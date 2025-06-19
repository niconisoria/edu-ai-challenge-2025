package validation

import (
	"testing"
)

func TestNumberValidator_ValidateValidNumbers(t *testing.T) {
	validator := &NumberValidator{}

	// Test with various valid numeric types
	testCases := []interface{}{
		float64(42.5),
		float32(42.5),
		int(42),
		int8(42),
		int16(42),
		int32(42),
		int64(42),
		uint(42),
		uint8(42),
		uint16(42),
		uint32(42),
		uint64(42),
		0,
		-1,
		1.5,
		-3.14,
	}

	for _, testCase := range testCases {
		result := validator.Validate(testCase)
		if !result.IsValid {
			t.Errorf("Number validator should accept valid number: %v (%T)", testCase, testCase)
		}
		if len(result.Errors) > 0 {
			t.Errorf("Number validator should not return errors for valid number: %v (%T)", testCase, testCase)
		}
	}
}

func TestNumberValidator_ValidateInvalidTypes(t *testing.T) {
	validator := &NumberValidator{}

	// Test with invalid types
	invalidTestCases := []interface{}{
		"42",
		"not a number",
		true,
		false,
		[]int{1, 2, 3},
		map[string]int{"key": 42},
		&NumberValidator{},
		"",
		"123.45",
	}

	for _, testCase := range invalidTestCases {
		result := validator.Validate(testCase)
		if result.IsValid {
			t.Errorf("Number validator should reject invalid type: %v (%T)", testCase, testCase)
		}
		if len(result.Errors) == 0 {
			t.Errorf("Number validator should return errors for invalid type: %v (%T)", testCase, testCase)
		}
		if result.Errors[0].Message == "" {
			t.Errorf("Number validator should return error message for invalid type: %v (%T)", testCase, testCase)
		}
	}
}

func TestNumberValidator_ValidateNil(t *testing.T) {
	validator := &NumberValidator{}

	// Test with nil value (not optional)
	result := validator.Validate(nil)
	if result.IsValid {
		t.Error("Number validator should reject nil value when not optional")
	}
	if len(result.Errors) == 0 {
		t.Error("Number validator should return error for nil value when not optional")
	}
	if result.Errors[0].Message == "" {
		t.Error("Number validator should return error message for nil value")
	}
}

func TestNumberValidator_ValidateNilOptional(t *testing.T) {
	validator := &NumberValidator{}

	// Test with nil value (optional)
	validator.Optional()
	result := validator.Validate(nil)
	if !result.IsValid {
		t.Error("Number validator should accept nil value when optional")
	}
	if len(result.Errors) > 0 {
		t.Error("Number validator should not return errors for nil value when optional")
	}
}

func TestNumberValidator_MinConstraint(t *testing.T) {
	validator := &NumberValidator{}
	minValue := 10.0

	validator.Min(minValue)

	// Test with value below min
	result := validator.Validate(5.0)
	if result.IsValid {
		t.Error("Number validator should reject value below min")
	}
	if len(result.Errors) == 0 {
		t.Error("Number validator should return error for value below min")
	}
	if result.Errors[0].Message == "" {
		t.Error("Number validator should return error message for value below min")
	}

	// Test with value equal to min
	result = validator.Validate(10.0)
	if !result.IsValid {
		t.Error("Number validator should accept value equal to min")
	}

	// Test with value above min
	result = validator.Validate(15.0)
	if !result.IsValid {
		t.Error("Number validator should accept value above min")
	}
}

func TestNumberValidator_MaxConstraint(t *testing.T) {
	validator := &NumberValidator{}
	maxValue := 100.0

	validator.Max(maxValue)

	// Test with value above max
	result := validator.Validate(150.0)
	if result.IsValid {
		t.Error("Number validator should reject value above max")
	}
	if len(result.Errors) == 0 {
		t.Error("Number validator should return error for value above max")
	}
	if result.Errors[0].Message == "" {
		t.Error("Number validator should return error message for value above max")
	}

	// Test with value equal to max
	result = validator.Validate(100.0)
	if !result.IsValid {
		t.Error("Number validator should accept value equal to max")
	}

	// Test with value below max
	result = validator.Validate(50.0)
	if !result.IsValid {
		t.Error("Number validator should accept value below max")
	}
}

func TestNumberValidator_MinMaxConstraints(t *testing.T) {
	validator := &NumberValidator{}
	minValue := 10.0
	maxValue := 100.0

	validator.Min(minValue).Max(maxValue)

	// Test with value below min
	result := validator.Validate(5.0)
	if result.IsValid {
		t.Error("Number validator should reject value below min when both min and max are set")
	}

	// Test with value above max
	result = validator.Validate(150.0)
	if result.IsValid {
		t.Error("Number validator should reject value above max when both min and max are set")
	}

	// Test with value in range
	result = validator.Validate(50.0)
	if !result.IsValid {
		t.Error("Number validator should accept value within min-max range")
	}

	// Test with value equal to min
	result = validator.Validate(10.0)
	if !result.IsValid {
		t.Error("Number validator should accept value equal to min")
	}

	// Test with value equal to max
	result = validator.Validate(100.0)
	if !result.IsValid {
		t.Error("Number validator should accept value equal to max")
	}
}

func TestNumberValidator_Min(t *testing.T) {
	validator := &NumberValidator{}
	minValue := 10.5

	validator.Min(minValue)

	if validator.min == nil {
		t.Error("Min should set the min field")
	}
	if *validator.min != minValue {
		t.Errorf("Expected min %f, got %f", minValue, *validator.min)
	}
}

func TestNumberValidator_Max(t *testing.T) {
	validator := &NumberValidator{}
	maxValue := 100.0

	validator.Max(maxValue)

	if validator.max == nil {
		t.Error("Max should set the max field")
	}
	if *validator.max != maxValue {
		t.Errorf("Expected max %f, got %f", maxValue, *validator.max)
	}
}

func TestNumberValidator_Optional(t *testing.T) {
	validator := &NumberValidator{}

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

func TestNumberValidator_WithMessage(t *testing.T) {
	validator := &NumberValidator{}
	customMessage := "Custom number validation message"

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

func TestNumberValidator_WithMessageValidation(t *testing.T) {
	validator := &NumberValidator{}
	customMessage := "Custom number validation message"

	validator.WithMessage(customMessage)

	// Test that custom message is used for invalid type
	result := validator.Validate("not a number")
	if result.IsValid {
		t.Error("Number validator should reject string value")
	}
	if len(result.Errors) == 0 {
		t.Error("Number validator should return error for string value")
	}
	if result.Errors[0].Message != customMessage {
		t.Errorf("Expected custom message '%s', got '%s'", customMessage, result.Errors[0].Message)
	}
}

func TestNumberValidator_Chaining(t *testing.T) {
	validator := &NumberValidator{}

	// Test method chaining
	result := validator.
		Min(0.0).
		Max(100.0).
		Optional().
		WithMessage("Custom message")

	// Should return the same validator
	if result != validator {
		t.Error("Method chaining should return the same validator instance")
	}

	// Should have all properties set
	if validator.min == nil || *validator.min != 0.0 {
		t.Error("Min should be set to 0.0")
	}
	if validator.max == nil || *validator.max != 100.0 {
		t.Error("Max should be set to 100.0")
	}
	if !validator.isOptional() {
		t.Error("Validator should be optional")
	}
	if validator.getMessage("default") != "Custom message" {
		t.Error("Validator should have custom message")
	}
}

func TestNumberValidator_ZeroValues(t *testing.T) {
	validator := &NumberValidator{}

	// Test with zero values
	validator.Min(0.0)
	validator.Max(0.0)

	if validator.min == nil || *validator.min != 0.0 {
		t.Error("Min should accept zero value")
	}
	if validator.max == nil || *validator.max != 0.0 {
		t.Error("Max should accept zero value")
	}

	// Test validation with zero constraints
	result := validator.Validate(0.0)
	if !result.IsValid {
		t.Error("Number validator should accept zero value when min and max are both zero")
	}

	result = validator.Validate(1.0)
	if result.IsValid {
		t.Error("Number validator should reject positive value when max is zero")
	}

	result = validator.Validate(-1.0)
	if result.IsValid {
		t.Error("Number validator should reject negative value when min is zero")
	}
}

func TestNumberValidator_NegativeValues(t *testing.T) {
	validator := &NumberValidator{}

	// Test with negative values
	validator.Min(-100.0)
	validator.Max(-1.0)

	if validator.min == nil || *validator.min != -100.0 {
		t.Error("Min should accept negative value")
	}
	if validator.max == nil || *validator.max != -1.0 {
		t.Error("Max should accept negative value")
	}

	// Test validation with negative constraints
	result := validator.Validate(-50.0)
	if !result.IsValid {
		t.Error("Number validator should accept value within negative range")
	}

	result = validator.Validate(-150.0)
	if result.IsValid {
		t.Error("Number validator should reject value below negative min")
	}

	result = validator.Validate(0.0)
	if result.IsValid {
		t.Error("Number validator should reject positive value when max is negative")
	}
}

func TestNumberValidator_InterfaceCompliance(t *testing.T) {
	validator := &NumberValidator{}

	// Test that it implements the Validator interface
	var _ Validator[float64] = validator
}

func TestNumberValidator_DefaultState(t *testing.T) {
	validator := &NumberValidator{}

	// Test default state
	if validator.isOptional() {
		t.Error("New NumberValidator should not be optional by default")
	}

	// Test default message behavior
	defaultMsg := "Default message"
	if validator.getMessage(defaultMsg) != defaultMsg {
		t.Error("New NumberValidator should return default message when no custom message is set")
	}

	// Test that min and max are nil by default
	if validator.min != nil {
		t.Error("New NumberValidator should have nil min by default")
	}
	if validator.max != nil {
		t.Error("New NumberValidator should have nil max by default")
	}
}

func TestNumberValidator_ErrorMessages(t *testing.T) {
	validator := &NumberValidator{}

	// Test nil value error message
	result := validator.Validate(nil)
	if result.IsValid {
		t.Error("Number validator should reject nil value")
	}
	if len(result.Errors) == 0 {
		t.Error("Number validator should return error for nil value")
	}
	expectedNilMsg := "Number value is required"
	if result.Errors[0].Message != expectedNilMsg {
		t.Errorf("Expected nil error message '%s', got '%s'", expectedNilMsg, result.Errors[0].Message)
	}

	// Test invalid type error message
	result = validator.Validate("string value")
	if result.IsValid {
		t.Error("Number validator should reject string value")
	}
	if len(result.Errors) == 0 {
		t.Error("Number validator should return error for string value")
	}
	expectedTypeMsg := "Expected numeric value, got string"
	if result.Errors[0].Message != expectedTypeMsg {
		t.Errorf("Expected type error message '%s', got '%s'", expectedTypeMsg, result.Errors[0].Message)
	}

	// Test min constraint error message
	validator.Min(10.0)
	result = validator.Validate(5.0)
	if result.IsValid {
		t.Error("Number validator should reject value below min")
	}
	if len(result.Errors) == 0 {
		t.Error("Number validator should return error for value below min")
	}
	expectedMinMsg := "Number must be at least 10.000000"
	if result.Errors[0].Message != expectedMinMsg {
		t.Errorf("Expected min error message '%s', got '%s'", expectedMinMsg, result.Errors[0].Message)
	}

	// Test max constraint error message
	validator = &NumberValidator{}
	validator.Max(100.0)
	result = validator.Validate(150.0)
	if result.IsValid {
		t.Error("Number validator should reject value above max")
	}
	if len(result.Errors) == 0 {
		t.Error("Number validator should return error for value above max")
	}
	expectedMaxMsg := "Number must be at most 100.000000"
	if result.Errors[0].Message != expectedMaxMsg {
		t.Errorf("Expected max error message '%s', got '%s'", expectedMaxMsg, result.Errors[0].Message)
	}
}

func TestNumberValidator_DifferentNumericTypes(t *testing.T) {
	validator := &NumberValidator{}
	validator.Min(10.0).Max(100.0)

	// Test with different numeric types within range
	testCases := []interface{}{
		float64(50.0),
		float32(50.0),
		int(50),
		int8(50),
		int16(50),
		int32(50),
		int64(50),
		uint(50),
		uint8(50),
		uint16(50),
		uint32(50),
		uint64(50),
	}

	for _, testCase := range testCases {
		result := validator.Validate(testCase)
		if !result.IsValid {
			t.Errorf("Number validator should accept valid numeric type %T: %v", testCase, testCase)
		}
	}

	// Test with different numeric types outside range
	outOfRangeCases := []interface{}{
		float64(5.0),
		float32(5.0),
		int(5),
		int8(5),
		int16(5),
		int32(5),
		int64(5),
		uint(5),
		uint8(5),
		uint16(5),
		uint32(5),
		uint64(5),
	}

	for _, testCase := range outOfRangeCases {
		result := validator.Validate(testCase)
		if result.IsValid {
			t.Errorf("Number validator should reject out-of-range numeric type %T: %v", testCase, testCase)
		}
	}
}
