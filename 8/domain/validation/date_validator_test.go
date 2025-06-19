package validation

import (
	"testing"
	"time"
)

func TestDateValidator_Validate(t *testing.T) {
	validator := &DateValidator{}

	// Test with valid time values
	now := time.Now()
	testCases := []interface{}{
		now,
		time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(1990, 12, 31, 23, 59, 59, 999999999, time.UTC),
	}

	for _, testCase := range testCases {
		result := validator.Validate(testCase)
		if !result.IsValid {
			t.Errorf("Date validator should accept valid time value: %v", testCase)
		}
		if len(result.Errors) > 0 {
			t.Errorf("Date validator should not return errors for valid time value: %v", testCase)
		}
	}
}

func TestDateValidator_ValidateInvalidTypes(t *testing.T) {
	validator := &DateValidator{}

	// Test with invalid types
	invalidTestCases := []interface{}{
		"2023-01-01",
		"2023/01/01",
		"01/01/2023",
		"date",
		"time",
		1,
		0,
		-1,
		1.5,
		true,
		false,
		[]time.Time{time.Now()},
		map[string]time.Time{"key": time.Now()},
		&DateValidator{},
	}

	for _, testCase := range invalidTestCases {
		result := validator.Validate(testCase)
		if result.IsValid {
			t.Errorf("Date validator should reject invalid type: %v (%T)", testCase, testCase)
		}
		if len(result.Errors) == 0 {
			t.Errorf("Date validator should return errors for invalid type: %v (%T)", testCase, testCase)
		}
		if result.Errors[0].Message == "" {
			t.Errorf("Date validator should return error message for invalid type: %v (%T)", testCase, testCase)
		}
	}
}

func TestDateValidator_ValidateNil(t *testing.T) {
	validator := &DateValidator{}

	// Test with nil value (not optional)
	result := validator.Validate(nil)
	if result.IsValid {
		t.Error("Date validator should reject nil value when not optional")
	}
	if len(result.Errors) == 0 {
		t.Error("Date validator should return error for nil value when not optional")
	}
	if result.Errors[0].Message == "" {
		t.Error("Date validator should return error message for nil value")
	}
}

func TestDateValidator_ValidateNilOptional(t *testing.T) {
	validator := &DateValidator{}

	// Test with nil value (optional)
	validator.Optional()
	result := validator.Validate(nil)
	if !result.IsValid {
		t.Error("Date validator should accept nil value when optional")
	}
	if len(result.Errors) > 0 {
		t.Error("Date validator should not return errors for nil value when optional")
	}
}

func TestDateValidator_Optional(t *testing.T) {
	validator := &DateValidator{}

	result := validator.Optional()

	// Should return the same validator
	if result != validator {
		t.Error("Optional() should return the same validator instance")
	}

	// Should be marked as optional
	if !validator.isOptional() {
		t.Error("Validator should be marked as optional after Optional() call")
	}
}

func TestDateValidator_WithMessage(t *testing.T) {
	validator := &DateValidator{}
	customMessage := "Custom date validation message"

	result := validator.WithMessage(customMessage)

	// Should return the same validator
	if result != validator {
		t.Error("WithMessage() should return the same validator instance")
	}

	// Should have the custom message
	if validator.getMessage("default") != customMessage {
		t.Error("Validator should have the custom message after WithMessage() call")
	}
}

func TestDateValidator_WithMessageValidation(t *testing.T) {
	validator := &DateValidator{}
	customMessage := "Custom date validation message"

	validator.WithMessage(customMessage)

	// Test that custom message is used for invalid type
	result := validator.Validate("not a date")
	if result.IsValid {
		t.Error("Date validator should reject string value")
	}
	if len(result.Errors) == 0 {
		t.Error("Date validator should return error for string value")
	}
	if result.Errors[0].Message != customMessage {
		t.Errorf("Expected custom message '%s', got '%s'", customMessage, result.Errors[0].Message)
	}
}

func TestDateValidator_Chaining(t *testing.T) {
	validator := &DateValidator{}

	// Test method chaining
	result := validator.
		Optional().
		WithMessage("Custom message")

	// Should return the same validator
	if result != validator {
		t.Error("Method chaining should return the same validator instance")
	}

	// Should have all properties set
	if !validator.isOptional() {
		t.Error("Validator should be optional")
	}
	if validator.getMessage("default") != "Custom message" {
		t.Error("Validator should have custom message")
	}
}

func TestDateValidator_InterfaceCompliance(t *testing.T) {
	validator := &DateValidator{}

	// Test that it implements the Validator interface
	var _ Validator[time.Time] = validator
}

func TestDateValidator_DefaultState(t *testing.T) {
	validator := &DateValidator{}

	// Test default state
	if validator.isOptional() {
		t.Error("New DateValidator should not be optional by default")
	}

	// Test default message behavior
	defaultMsg := "Default message"
	if validator.getMessage(defaultMsg) != defaultMsg {
		t.Error("New DateValidator should return default message when no custom message is set")
	}
}

func TestDateValidator_EdgeCases(t *testing.T) {
	validator := &DateValidator{}

	// Test with zero time
	zeroTime := time.Time{}
	result := validator.Validate(zeroTime)
	if !result.IsValid {
		t.Error("Date validator should accept zero time")
	}

	// Test with very old date
	oldDate := time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
	result = validator.Validate(oldDate)
	if !result.IsValid {
		t.Error("Date validator should accept very old dates")
	}

	// Test with future date
	futureDate := time.Date(9999, 12, 31, 23, 59, 59, 999999999, time.UTC)
	result = validator.Validate(futureDate)
	if !result.IsValid {
		t.Error("Date validator should accept future dates")
	}
}

func TestDateValidator_ErrorMessages(t *testing.T) {
	validator := &DateValidator{}

	// Test nil value error message
	result := validator.Validate(nil)
	if result.IsValid {
		t.Error("Date validator should reject nil value")
	}
	if len(result.Errors) == 0 {
		t.Error("Date validator should return error for nil value")
	}
	expectedNilMsg := "Date value is required"
	if result.Errors[0].Message != expectedNilMsg {
		t.Errorf("Expected nil error message '%s', got '%s'", expectedNilMsg, result.Errors[0].Message)
	}

	// Test invalid type error message
	result = validator.Validate("string value")
	if result.IsValid {
		t.Error("Date validator should reject string value")
	}
	if len(result.Errors) == 0 {
		t.Error("Date validator should return error for string value")
	}
	expectedTypeMsg := "Expected time.Time value, got string"
	if result.Errors[0].Message != expectedTypeMsg {
		t.Errorf("Expected type error message '%s', got '%s'", expectedTypeMsg, result.Errors[0].Message)
	}
}

func TestDateValidator_DifferentTimeFormats(t *testing.T) {
	validator := &DateValidator{}

	// Test with different time formats and timezones
	testCases := []interface{}{
		time.Now().UTC(),
		time.Now().Local(),
		time.Date(2023, 1, 1, 12, 30, 45, 123456789, time.UTC),
		time.Date(2023, 12, 31, 23, 59, 59, 999999999, time.FixedZone("EST", -5*3600)),
		time.Date(2023, 6, 15, 6, 0, 0, 0, time.FixedZone("PST", -8*3600)),
	}

	for _, testCase := range testCases {
		result := validator.Validate(testCase)
		if !result.IsValid {
			t.Errorf("Date validator should accept time value: %v", testCase)
		}
		if len(result.Errors) > 0 {
			t.Errorf("Date validator should not return errors for time value: %v", testCase)
		}
	}
}
