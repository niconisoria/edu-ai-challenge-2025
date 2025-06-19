package validation

import (
	"testing"
	"time"
)

func TestIntegration_AllValidatorsTogether(t *testing.T) {
	schema := map[string]AnyValidator{
		"name":     (&StringValidator{}).MinLength(2).MaxLength(50),
		"age":      (&NumberValidator{}).Min(0).Max(150),
		"active":   &BooleanValidator{},
		"birthday": &DateValidator{},
		"tags":     &ArrayValidator[any]{ItemValidator: &StringValidator{}},
		"address": &ObjectValidator[map[string]any]{
			Schema: map[string]AnyValidator{
				"street": (&StringValidator{}).MinLength(3),
				"city":   &StringValidator{},
				"zip":    &StringValidator{},
			},
		},
	}

	validator := &ObjectValidator[map[string]any]{Schema: schema}

	// Valid object
	validObj := map[string]any{
		"name":     "John Doe",
		"age":      30,
		"active":   true,
		"birthday": time.Now(),
		"tags":     []any{"golang", "ai"},
		"address": map[string]any{
			"street": "123 Main St",
			"city":   "Metropolis",
			"zip":    "12345",
		},
	}
	result := validator.Validate(validObj)
	if !result.IsValid {
		t.Errorf("Expected valid object, got errors: %+v", result.Errors)
	}

	// Invalid object: multiple errors
	invalidObj := map[string]any{
		"name":     "A",              // too short
		"age":      200,              // too high
		"active":   "yes",            // not a boolean
		"birthday": "not-a-date",     // not a time.Time
		"tags":     []any{"ok", 123}, // second tag not a string
		"address": map[string]any{
			"street": "St", // too short
			// missing city
			"zip":   12345,   // not a string
			"extra": "field", // extra field
		},
		"extra_field": 42, // extra field at root
	}
	result = validator.Validate(invalidObj)
	if result.IsValid {
		t.Error("Expected invalid object, but got valid")
	}
	if len(result.Errors) < 7 {
		t.Errorf("Expected multiple errors, got %d: %+v", len(result.Errors), result.Errors)
	}

	// Check that each validator's error is present
	var (
		foundNameErr, foundAgeErr, foundActiveErr, foundBirthdayErr, foundTagsErr, foundAddressCityErr, foundAddressExtraErr, foundRootExtraErr bool
	)
	for _, err := range result.Errors {
		switch err.Field {
		case "name":
			foundNameErr = true
		case "age":
			foundAgeErr = true
		case "active":
			foundActiveErr = true
		case "birthday":
			foundBirthdayErr = true
		case "tags":
			foundTagsErr = true
		case "address":
			if err.Message == "Field 'city' is required" {
				foundAddressCityErr = true
			}
		case "extra_field":
			foundRootExtraErr = true
		}
		if err.Field == "address" && err.Message == "Unexpected field 'extra'" {
			foundAddressExtraErr = true
		}
	}
	if !foundNameErr {
		t.Error("Expected error for 'name' field")
	}
	if !foundAgeErr {
		t.Error("Expected error for 'age' field")
	}
	if !foundActiveErr {
		t.Error("Expected error for 'active' field")
	}
	if !foundBirthdayErr {
		t.Error("Expected error for 'birthday' field")
	}
	if !foundTagsErr {
		t.Error("Expected error for 'tags' field")
	}
	if !foundAddressCityErr {
		t.Error("Expected error for missing 'city' in address")
	}
	if !foundAddressExtraErr {
		t.Error("Expected error for extra field in address")
	}
	if !foundRootExtraErr {
		t.Error("Expected error for extra field at root level")
	}
}
