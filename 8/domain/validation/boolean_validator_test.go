package validation

import (
	"testing"
)

func TestBooleanValidator_Validate(t *testing.T) {
	validator := &BooleanValidator{}

	// Test with valid boolean values
	testCases := []interface{}{true, false}

	for _, testCase := range testCases {
		result := validator.Validate(testCase)
		if !result.IsValid {
			t.Errorf("Boolean validator should accept valid boolean value: %v", testCase)
		}
		if len(result.Errors) > 0 {
			t.Errorf("Boolean validator should not return errors for valid boolean value: %v", testCase)
		}
	}
}

func TestBooleanValidator_ValidateInvalidTypes(t *testing.T) {
	validator := &BooleanValidator{}

	// Test with invalid types
	invalidTestCases := []interface{}{
		"true",
		"false",
		"boolean",
		1,
		0,
		-1,
		1.5,
		[]bool{true, false},
		map[string]bool{"key": true},
		&BooleanValidator{},
	}

	for _, testCase := range invalidTestCases {
		result := validator.Validate(testCase)
		if result.IsValid {
			t.Errorf("Boolean validator should reject invalid type: %v (%T)", testCase, testCase)
		}
		if len(result.Errors) == 0 {
			t.Errorf("Boolean validator should return errors for invalid type: %v (%T)", testCase, testCase)
		}
		if result.Errors[0].Message == "" {
			t.Errorf("Boolean validator should return error message for invalid type: %v (%T)", testCase, testCase)
		}
	}
}

func TestBooleanValidator_ValidateNil(t *testing.T) {
	validator := &BooleanValidator{}

	// Test with nil value (not optional)
	result := validator.Validate(nil)
	if result.IsValid {
		t.Error("Boolean validator should reject nil value when not optional")
	}
	if len(result.Errors) == 0 {
		t.Error("Boolean validator should return error for nil value when not optional")
	}
	if result.Errors[0].Message == "" {
		t.Error("Boolean validator should return error message for nil value")
	}
}

func TestBooleanValidator_ValidateNilOptional(t *testing.T) {
	validator := &BooleanValidator{}

	// Test with nil value (optional)
	validator.Optional()
	result := validator.Validate(nil)
	if !result.IsValid {
		t.Error("Boolean validator should accept nil value when optional")
	}
	if len(result.Errors) > 0 {
		t.Error("Boolean validator should not return errors for nil value when optional")
	}
}

func TestBooleanValidator_Optional(t *testing.T) {
	validator := &BooleanValidator{}

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

func TestBooleanValidator_WithMessage(t *testing.T) {
	validator := &BooleanValidator{}
	customMessage := "Custom boolean validation message"

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

func TestBooleanValidator_WithMessageValidation(t *testing.T) {
	validator := &BooleanValidator{}
	customMessage := "Custom boolean validation message"

	validator.WithMessage(customMessage)

	// Test that custom message is used for invalid type
	result := validator.Validate("not a boolean")
	if result.IsValid {
		t.Error("Boolean validator should reject string value")
	}
	if len(result.Errors) == 0 {
		t.Error("Boolean validator should return error for string value")
	}
	if result.Errors[0].Message != customMessage {
		t.Errorf("Expected custom message '%s', got '%s'", customMessage, result.Errors[0].Message)
	}
}

func TestBooleanValidator_Chaining(t *testing.T) {
	validator := &BooleanValidator{}

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

func TestBooleanValidator_InterfaceCompliance(t *testing.T) {
	validator := &BooleanValidator{}

	// Test that it implements the Validator interface
	var _ Validator[bool] = validator
}

func TestBooleanValidator_DefaultState(t *testing.T) {
	validator := &BooleanValidator{}

	// Test default state
	if validator.isOptional() {
		t.Error("New BooleanValidator should not be optional by default")
	}

	// Test default message behavior
	defaultMsg := "Default message"
	if validator.getMessage(defaultMsg) != defaultMsg {
		t.Error("New BooleanValidator should return default message when no custom message is set")
	}
}

func TestBooleanValidator_ErrorMessages(t *testing.T) {
	validator := &BooleanValidator{}

	// Test nil value error message
	result := validator.Validate(nil)
	if result.IsValid {
		t.Error("Boolean validator should reject nil value")
	}
	if len(result.Errors) == 0 {
		t.Error("Boolean validator should return error for nil value")
	}
	expectedNilMsg := "Boolean value is required"
	if result.Errors[0].Message != expectedNilMsg {
		t.Errorf("Expected nil error message '%s', got '%s'", expectedNilMsg, result.Errors[0].Message)
	}

	// Test invalid type error message
	result = validator.Validate("string value")
	if result.IsValid {
		t.Error("Boolean validator should reject string value")
	}
	if len(result.Errors) == 0 {
		t.Error("Boolean validator should return error for string value")
	}
	expectedTypeMsg := "Expected boolean value, got string"
	if result.Errors[0].Message != expectedTypeMsg {
		t.Errorf("Expected type error message '%s', got '%s'", expectedTypeMsg, result.Errors[0].Message)
	}
}
