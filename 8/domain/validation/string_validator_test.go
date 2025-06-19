package validation

import (
	"testing"
)

func TestStringValidator_Validate(t *testing.T) {
	validator := &StringValidator{}

	// Test with valid string
	result := validator.Validate("hello")
	if !result.IsValid {
		t.Error("String validator should accept valid string")
	}
	if len(result.Errors) > 0 {
		t.Error("String validator should not return errors for valid string")
	}
}

func TestStringValidator_ValidateInvalidTypes(t *testing.T) {
	validator := &StringValidator{}

	// Test with invalid types
	invalidTestCases := []interface{}{
		42,
		3.14,
		true,
		false,
		[]string{"hello"},
		map[string]string{"key": "value"},
		&StringValidator{},
		[]int{1, 2, 3},
	}

	for _, testCase := range invalidTestCases {
		result := validator.Validate(testCase)
		if result.IsValid {
			t.Errorf("String validator should reject invalid type: %v (%T)", testCase, testCase)
		}
		if len(result.Errors) == 0 {
			t.Errorf("String validator should return errors for invalid type: %v (%T)", testCase, testCase)
		}
		if result.Errors[0].Message == "" {
			t.Errorf("String validator should return error message for invalid type: %v (%T)", testCase, testCase)
		}
	}
}

func TestStringValidator_ValidateNil(t *testing.T) {
	validator := &StringValidator{}

	// Test with nil value (not optional)
	result := validator.Validate(nil)
	if result.IsValid {
		t.Error("String validator should reject nil value when not optional")
	}
	if len(result.Errors) == 0 {
		t.Error("String validator should return error for nil value when not optional")
	}
}

func TestStringValidator_ValidateNilOptional(t *testing.T) {
	validator := &StringValidator{}

	// Test with nil value (optional)
	validator.Optional()
	result := validator.Validate(nil)
	if !result.IsValid {
		t.Error("String validator should accept nil value when optional")
	}
	if len(result.Errors) > 0 {
		t.Error("String validator should not return errors for nil value when optional")
	}
}

func TestStringValidator_MinLength(t *testing.T) {
	validator := &StringValidator{}
	minLength := 5

	validator.MinLength(minLength)

	if validator.minLength == nil {
		t.Error("MinLength should set the minLength field")
	}
	if *validator.minLength != minLength {
		t.Errorf("Expected minLength %d, got %d", minLength, *validator.minLength)
	}

	// Test validation with min length
	result := validator.Validate("hello") // 5 characters
	if !result.IsValid {
		t.Error("String validator should accept string with exact min length")
	}

	result = validator.Validate("helloworld") // 10 characters
	if !result.IsValid {
		t.Error("String validator should accept string longer than min length")
	}

	result = validator.Validate("hi") // 2 characters
	if result.IsValid {
		t.Error("String validator should reject string shorter than min length")
	}
	if len(result.Errors) == 0 {
		t.Error("String validator should return error for string shorter than min length")
	}
}

func TestStringValidator_MaxLength(t *testing.T) {
	validator := &StringValidator{}
	maxLength := 10

	validator.MaxLength(maxLength)

	if validator.maxLength == nil {
		t.Error("MaxLength should set the maxLength field")
	}
	if *validator.maxLength != maxLength {
		t.Errorf("Expected maxLength %d, got %d", maxLength, *validator.maxLength)
	}

	// Test validation with max length
	result := validator.Validate("hello") // 5 characters
	if !result.IsValid {
		t.Error("String validator should accept string shorter than max length")
	}

	result = validator.Validate("helloworld") // 10 characters
	if !result.IsValid {
		t.Error("String validator should accept string with exact max length")
	}

	result = validator.Validate("helloworld123") // 13 characters
	if result.IsValid {
		t.Error("String validator should reject string longer than max length")
	}
	if len(result.Errors) == 0 {
		t.Error("String validator should return error for string longer than max length")
	}
}

func TestStringValidator_Pattern(t *testing.T) {
	validator := &StringValidator{}
	pattern := `^[a-z]+$`

	validator.Pattern(pattern)

	if validator.pattern == nil {
		t.Error("Pattern should set the pattern field")
	}

	// Test that the pattern was compiled correctly
	if validator.pattern.String() != pattern {
		t.Errorf("Expected pattern '%s', got '%s'", pattern, validator.pattern.String())
	}

	// Test validation with pattern
	result := validator.Validate("hello") // matches pattern
	if !result.IsValid {
		t.Error("String validator should accept string matching pattern")
	}

	result = validator.Validate("hello123") // doesn't match pattern
	if result.IsValid {
		t.Error("String validator should reject string not matching pattern")
	}
	if len(result.Errors) == 0 {
		t.Error("String validator should return error for string not matching pattern")
	}

	result = validator.Validate("HELLO") // doesn't match pattern
	if result.IsValid {
		t.Error("String validator should reject string not matching pattern")
	}
}

func TestStringValidator_CombinedConstraints(t *testing.T) {
	validator := &StringValidator{}
	validator.MinLength(3).MaxLength(10).Pattern(`^[a-z]+$`)

	// Test valid string
	result := validator.Validate("hello") // 5 chars, lowercase
	if !result.IsValid {
		t.Error("String validator should accept string meeting all constraints")
	}

	// Test too short
	result = validator.Validate("hi") // 2 chars
	if result.IsValid {
		t.Error("String validator should reject string too short")
	}

	// Test too long
	result = validator.Validate("helloworld123") // 13 chars
	if result.IsValid {
		t.Error("String validator should reject string too long")
	}

	// Test wrong pattern
	result = validator.Validate("HELLO") // uppercase
	if result.IsValid {
		t.Error("String validator should reject string not matching pattern")
	}
}

func TestStringValidator_EmptyString(t *testing.T) {
	validator := &StringValidator{}

	// Test empty string without constraints
	result := validator.Validate("")
	if !result.IsValid {
		t.Error("String validator should accept empty string without constraints")
	}

	// Test empty string with min length constraint
	validator.MinLength(1)
	result = validator.Validate("")
	if result.IsValid {
		t.Error("String validator should reject empty string when min length > 0")
	}
}

func TestStringValidator_Optional(t *testing.T) {
	validator := &StringValidator{}

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

func TestStringValidator_WithMessage(t *testing.T) {
	validator := &StringValidator{}
	customMessage := "Custom string validation message"

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

func TestStringValidator_WithMessageValidation(t *testing.T) {
	validator := &StringValidator{}
	customMessage := "Custom string validation message"

	validator.WithMessage(customMessage)

	// Test that custom message is used for invalid type
	result := validator.Validate(42)
	if result.IsValid {
		t.Error("String validator should reject numeric value")
	}
	if len(result.Errors) == 0 {
		t.Error("String validator should return error for numeric value")
	}
	if result.Errors[0].Message != customMessage {
		t.Errorf("Expected custom message '%s', got '%s'", customMessage, result.Errors[0].Message)
	}
}

func TestStringValidator_Chaining(t *testing.T) {
	validator := &StringValidator{}

	// Test method chaining
	result := validator.
		MinLength(3).
		MaxLength(10).
		Pattern(`^[a-z]+$`).
		Optional().
		WithMessage("Custom message")

	// Should return the same validator
	if result != validator {
		t.Error("Method chaining should return the same validator instance")
	}

	// Should have all properties set
	if validator.minLength == nil || *validator.minLength != 3 {
		t.Error("MinLength should be set to 3")
	}
	if validator.maxLength == nil || *validator.maxLength != 10 {
		t.Error("MaxLength should be set to 10")
	}
	if validator.pattern == nil {
		t.Error("Pattern should be set")
	}
	if !validator.isOptional() {
		t.Error("Validator should be optional")
	}
	if validator.getMessage("default") != "Custom message" {
		t.Error("Validator should have custom message")
	}
}

func TestStringValidator_InterfaceCompliance(t *testing.T) {
	validator := &StringValidator{}

	// Test that it implements the Validator interface
	var _ Validator[string] = validator
}

func TestStringValidator_DefaultState(t *testing.T) {
	validator := &StringValidator{}

	// Test default state
	if validator.isOptional() {
		t.Error("New StringValidator should not be optional by default")
	}

	// Test default message behavior
	defaultMsg := "Default message"
	if validator.getMessage(defaultMsg) != defaultMsg {
		t.Error("New StringValidator should return default message when no custom message is set")
	}

	// Test that constraints are nil by default
	if validator.minLength != nil {
		t.Error("New StringValidator should have nil minLength by default")
	}
	if validator.maxLength != nil {
		t.Error("New StringValidator should have nil maxLength by default")
	}
	if validator.pattern != nil {
		t.Error("New StringValidator should have nil pattern by default")
	}
}

func TestStringValidator_ErrorMessages(t *testing.T) {
	validator := &StringValidator{}

	// Test nil value error message
	result := validator.Validate(nil)
	if result.IsValid {
		t.Error("String validator should reject nil value")
	}
	if len(result.Errors) == 0 {
		t.Error("String validator should return error for nil value")
	}
	expectedNilMsg := "String value is required"
	if result.Errors[0].Message != expectedNilMsg {
		t.Errorf("Expected nil error message '%s', got '%s'", expectedNilMsg, result.Errors[0].Message)
	}

	// Test invalid type error message
	result = validator.Validate(42)
	if result.IsValid {
		t.Error("String validator should reject numeric value")
	}
	if len(result.Errors) == 0 {
		t.Error("String validator should return error for numeric value")
	}
	expectedTypeMsg := "Expected string value, got int"
	if result.Errors[0].Message != expectedTypeMsg {
		t.Errorf("Expected type error message '%s', got '%s'", expectedTypeMsg, result.Errors[0].Message)
	}

	// Test min length error message
	validator.MinLength(5)
	result = validator.Validate("hi")
	if result.IsValid {
		t.Error("String validator should reject short string")
	}
	if len(result.Errors) == 0 {
		t.Error("String validator should return error for short string")
	}
	expectedMinMsg := "String must be at least 5 characters long"
	if result.Errors[0].Message != expectedMinMsg {
		t.Errorf("Expected min length error message '%s', got '%s'", expectedMinMsg, result.Errors[0].Message)
	}

	// Test max length error message
	validator.MaxLength(10)
	result = validator.Validate("helloworld123")
	if result.IsValid {
		t.Error("String validator should reject long string")
	}
	if len(result.Errors) == 0 {
		t.Error("String validator should return error for long string")
	}
	expectedMaxMsg := "String must be at most 10 characters long"
	if result.Errors[0].Message != expectedMaxMsg {
		t.Errorf("Expected max length error message '%s', got '%s'", expectedMaxMsg, result.Errors[0].Message)
	}

	// Test pattern error message
	validator.Pattern(`^[a-z]+$`)
	result = validator.Validate("HELLO")
	if result.IsValid {
		t.Error("String validator should reject string not matching pattern")
	}
	if len(result.Errors) == 0 {
		t.Error("String validator should return error for string not matching pattern")
	}
	expectedPatternMsg := "String must match pattern: ^[a-z]+$"
	if result.Errors[0].Message != expectedPatternMsg {
		t.Errorf("Expected pattern error message '%s', got '%s'", expectedPatternMsg, result.Errors[0].Message)
	}
}

func TestStringValidator_ComplexPatterns(t *testing.T) {
	validator := &StringValidator{}

	// Test email pattern
	validator.Pattern(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	validEmails := []string{
		"test@example.com",
		"user.name@domain.co.uk",
		"user+tag@example.org",
	}

	for _, email := range validEmails {
		result := validator.Validate(email)
		if !result.IsValid {
			t.Errorf("String validator should accept valid email: %s", email)
		}
	}

	invalidEmails := []string{
		"invalid-email",
		"@example.com",
		"user@",
		"user@.com",
	}

	for _, email := range invalidEmails {
		result := validator.Validate(email)
		if result.IsValid {
			t.Errorf("String validator should reject invalid email: %s", email)
		}
	}
}

func TestStringValidator_EdgeCases(t *testing.T) {
	validator := &StringValidator{}

	// Test very long string
	longString := string(make([]byte, 10000))
	result := validator.Validate(longString)
	if !result.IsValid {
		t.Error("String validator should accept very long string without constraints")
	}

	// Test string with special characters
	specialString := "!@#$%^&*()_+-=[]{}|;':\",./<>?"
	result = validator.Validate(specialString)
	if !result.IsValid {
		t.Error("String validator should accept string with special characters")
	}

	// Test unicode string
	unicodeString := "Hello 世界 🌍"
	result = validator.Validate(unicodeString)
	if !result.IsValid {
		t.Error("String validator should accept unicode string")
	}
}
