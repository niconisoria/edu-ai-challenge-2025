package validation

import (
	"testing"
)

func TestValidationError_Error(t *testing.T) {
	// Test with field and message
	error := ValidationError{
		Field:   "email",
		Message: "Invalid email format",
	}

	expected := "email: Invalid email format"
	if error.Error() != expected {
		t.Errorf("Expected error message '%s', got '%s'", expected, error.Error())
	}
}

func TestValidationError_ErrorWithoutField(t *testing.T) {
	// Test with only message (no field)
	error := ValidationError{
		Field:   "",
		Message: "General validation error",
	}

	expected := "General validation error"
	if error.Error() != expected {
		t.Errorf("Expected error message '%s', got '%s'", expected, error.Error())
	}
}

func TestValidationResult_IsValid(t *testing.T) {
	// Test valid result
	validResult := ValidationResult{
		IsValid: true,
		Errors:  []ValidationError{},
	}

	if !validResult.IsValid {
		t.Error("ValidationResult should be valid when IsValid is true")
	}

	// Test invalid result
	invalidResult := ValidationResult{
		IsValid: false,
		Errors: []ValidationError{
			{Field: "email", Message: "Invalid email"},
		},
	}

	if invalidResult.IsValid {
		t.Error("ValidationResult should be invalid when IsValid is false")
	}
}

func TestValidationResult_Errors(t *testing.T) {
	// Test with multiple errors
	errors := []ValidationError{
		{Field: "email", Message: "Invalid email format"},
		{Field: "password", Message: "Password too short"},
		{Field: "age", Message: "Age must be positive"},
	}

	result := ValidationResult{
		IsValid: false,
		Errors:  errors,
	}

	if len(result.Errors) != 3 {
		t.Errorf("Expected 3 errors, got %d", len(result.Errors))
	}

	// Test error messages
	expectedMessages := []string{
		"email: Invalid email format",
		"password: Password too short",
		"age: Age must be positive",
	}

	for i, expected := range expectedMessages {
		if result.Errors[i].Error() != expected {
			t.Errorf("Expected error message '%s', got '%s'", expected, result.Errors[i].Error())
		}
	}
}

func TestValidationResult_EmptyErrors(t *testing.T) {
	// Test with no errors
	result := ValidationResult{
		IsValid: true,
		Errors:  []ValidationError{},
	}

	if len(result.Errors) != 0 {
		t.Errorf("Expected 0 errors, got %d", len(result.Errors))
	}
}

func TestValidationResult_NilErrors(t *testing.T) {
	// Test with nil errors slice
	result := ValidationResult{
		IsValid: true,
		Errors:  nil,
	}

	if result.Errors != nil {
		t.Error("Errors should be nil")
	}
}

func TestValidationError_FieldAccess(t *testing.T) {
	error := ValidationError{
		Field:   "username",
		Message: "Username is required",
	}

	if error.Field != "username" {
		t.Errorf("Expected field 'username', got '%s'", error.Field)
	}

	if error.Message != "Username is required" {
		t.Errorf("Expected message 'Username is required', got '%s'", error.Message)
	}
}

func TestValidationResult_Consistency(t *testing.T) {
	// Test that IsValid and Errors are consistent
	// When IsValid is true, there should be no errors
	validResult := ValidationResult{
		IsValid: true,
		Errors:  []ValidationError{},
	}

	if !validResult.IsValid || len(validResult.Errors) > 0 {
		t.Error("Valid result should have IsValid=true and no errors")
	}

	// When there are errors, IsValid should be false
	invalidResult := ValidationResult{
		IsValid: false,
		Errors: []ValidationError{
			{Field: "test", Message: "Test error"},
		},
	}

	if invalidResult.IsValid || len(invalidResult.Errors) == 0 {
		t.Error("Invalid result should have IsValid=false and at least one error")
	}
}
