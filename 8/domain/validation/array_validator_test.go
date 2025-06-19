package validation

import (
	"testing"
)

func TestArrayValidator_Validate(t *testing.T) {
	validator := &ArrayValidator[any]{}

	// Test with valid arrays
	testCases := []interface{}{
		[]interface{}{"item1", "item2", "item3"},
		[]interface{}{1, 2, 3},
		[]interface{}{true, false},
		[]interface{}{},
		[]string{"a", "b", "c"},
		[]int{1, 2, 3},
		[]bool{true, false},
		[]float64{1.1, 2.2, 3.3},
	}

	for _, testCase := range testCases {
		result := validator.Validate(testCase)
		if !result.IsValid {
			t.Errorf("Array validator should accept valid array: %v", testCase)
		}
		if len(result.Errors) > 0 {
			t.Errorf("Array validator should not return errors for valid array: %v", testCase)
		}
	}
}

func TestArrayValidator_ValidateInvalidTypes(t *testing.T) {
	validator := &ArrayValidator[any]{}

	// Test with invalid types
	invalidTestCases := []interface{}{
		"not an array",
		"",
		42,
		3.14,
		true,
		false,
		map[string]interface{}{"key": "value"},
		&ArrayValidator[any]{},
		nil, // will be tested separately for optional behavior
	}

	for _, testCase := range invalidTestCases {
		result := validator.Validate(testCase)
		if result.IsValid {
			t.Errorf("Array validator should reject invalid type: %v (%T)", testCase, testCase)
		}
		if len(result.Errors) == 0 {
			t.Errorf("Array validator should return errors for invalid type: %v (%T)", testCase, testCase)
		}
		if result.Errors[0].Message == "" {
			t.Errorf("Array validator should return error message for invalid type: %v (%T)", testCase, testCase)
		}
	}
}

func TestArrayValidator_ValidateNil(t *testing.T) {
	validator := &ArrayValidator[any]{}

	// Test with nil value (not optional)
	result := validator.Validate(nil)
	if result.IsValid {
		t.Error("Array validator should reject nil value when not optional")
	}
	if len(result.Errors) == 0 {
		t.Error("Array validator should return error for nil value when not optional")
	}
	if result.Errors[0].Message == "" {
		t.Error("Array validator should return error message for nil value")
	}
}

func TestArrayValidator_ValidateNilOptional(t *testing.T) {
	validator := &ArrayValidator[any]{}

	// Test with nil value (optional)
	validator.Optional()
	result := validator.Validate(nil)
	if !result.IsValid {
		t.Error("Array validator should accept nil value when optional")
	}
	if len(result.Errors) > 0 {
		t.Error("Array validator should not return errors for nil value when optional")
	}
}

func TestArrayValidator_WithItemValidator(t *testing.T) {
	// Create an item validator
	itemValidator := &StringValidator{}

	validator := &ArrayValidator[any]{
		ItemValidator: itemValidator,
	}

	// Test that item validator is set correctly
	if validator.ItemValidator != itemValidator {
		t.Error("ItemValidator should be set correctly")
	}

	// Test with valid string array
	validArray := []interface{}{"item1", "item2", "item3"}
	result := validator.Validate(validArray)
	if !result.IsValid {
		t.Error("Array validator should accept valid string array")
	}

	// Test with invalid string array (contains non-string items)
	invalidArray := []interface{}{"item1", 42, "item3"}
	result = validator.Validate(invalidArray)
	if result.IsValid {
		t.Error("Array validator should reject array with invalid items")
	}
	if len(result.Errors) == 0 {
		t.Error("Array validator should return errors for array with invalid items")
	}
	// Check that error is reported for the specific invalid item
	if len(result.Errors) > 0 && result.Errors[0].Field != "[1]" {
		t.Errorf("Expected error field '[1]', got '%s'", result.Errors[0].Field)
	}
}

func TestArrayValidator_WithItemValidatorComplex(t *testing.T) {
	// Create a number validator with constraints
	itemValidator := &NumberValidator{}
	itemValidator.Min(0).Max(100)

	validator := &ArrayValidator[any]{
		ItemValidator: itemValidator,
	}

	// Test with valid number array
	validArray := []interface{}{10, 50, 99}
	result := validator.Validate(validArray)
	if !result.IsValid {
		t.Error("Array validator should accept valid number array")
	}

	// Test with invalid number array (contains out-of-range numbers)
	invalidArray := []interface{}{10, 150, 99}
	result = validator.Validate(invalidArray)
	if result.IsValid {
		t.Error("Array validator should reject array with invalid numbers")
	}
	if len(result.Errors) == 0 {
		t.Error("Array validator should return errors for array with invalid numbers")
	}
	// Check that error is reported for the specific invalid item
	if len(result.Errors) > 0 && result.Errors[0].Field != "[1]" {
		t.Errorf("Expected error field '[1]', got '%s'", result.Errors[0].Field)
	}
}

func TestArrayValidator_WithItemValidatorMultipleErrors(t *testing.T) {
	// Create a string validator with length constraints
	itemValidator := &StringValidator{}
	itemValidator.MinLength(3).MaxLength(10)

	validator := &ArrayValidator[any]{
		ItemValidator: itemValidator,
	}

	// Test with array containing multiple invalid items
	invalidArray := []interface{}{"ab", "valid_string", "too_long_string_for_validation"}
	result := validator.Validate(invalidArray)
	if result.IsValid {
		t.Error("Array validator should reject array with multiple invalid items")
	}
	if len(result.Errors) < 2 {
		t.Errorf("Expected at least 2 errors, got %d", len(result.Errors))
	}

	// Check that errors are reported for the correct indices
	errorFields := make(map[string]bool)
	for _, err := range result.Errors {
		errorFields[err.Field] = true
	}

	if !errorFields["[0]"] {
		t.Error("Expected error for index [0]")
	}
	if !errorFields["[2]"] {
		t.Error("Expected error for index [2]")
	}
}

func TestArrayValidator_EmptyArray(t *testing.T) {
	validator := &ArrayValidator[any]{}

	// Test with empty array
	emptyArray := []interface{}{}

	result := validator.Validate(emptyArray)
	if !result.IsValid {
		t.Error("Array validator should accept empty array")
	}
}

func TestArrayValidator_EmptyArrayWithItemValidator(t *testing.T) {
	// Create an item validator
	itemValidator := &StringValidator{}

	validator := &ArrayValidator[any]{
		ItemValidator: itemValidator,
	}

	// Test with empty array (should be valid even with item validator)
	emptyArray := []interface{}{}

	result := validator.Validate(emptyArray)
	if !result.IsValid {
		t.Error("Array validator should accept empty array even with item validator")
	}
}

func TestArrayValidator_Optional(t *testing.T) {
	validator := &ArrayValidator[any]{}

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

func TestArrayValidator_WithMessage(t *testing.T) {
	validator := &ArrayValidator[any]{}
	customMessage := "Custom array validation message"

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

func TestArrayValidator_WithMessageValidation(t *testing.T) {
	validator := &ArrayValidator[any]{}
	customMessage := "Custom array validation message"

	validator.WithMessage(customMessage)

	// Test that custom message is used for invalid type
	result := validator.Validate("not an array")
	if result.IsValid {
		t.Error("Array validator should reject string value")
	}
	if len(result.Errors) == 0 {
		t.Error("Array validator should return error for string value")
	}
	if result.Errors[0].Message != customMessage {
		t.Errorf("Expected custom message '%s', got '%s'", customMessage, result.Errors[0].Message)
	}
}

func TestArrayValidator_Chaining(t *testing.T) {
	validator := &ArrayValidator[any]{}

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

func TestArrayValidator_InterfaceCompliance(t *testing.T) {
	validator := &ArrayValidator[any]{}

	// Test that it implements the Validator interface
	var _ Validator[any] = validator
}

func TestArrayValidator_DefaultState(t *testing.T) {
	validator := &ArrayValidator[any]{}

	// Test default state
	if validator.isOptional() {
		t.Error("New ArrayValidator should not be optional by default")
	}

	// Test default message behavior
	defaultMsg := "Default message"
	if validator.getMessage(defaultMsg) != defaultMsg {
		t.Error("New ArrayValidator should return default message when no custom message is set")
	}

	// Test that item validator is nil by default
	if validator.ItemValidator != nil {
		t.Error("New ArrayValidator should have nil ItemValidator by default")
	}
}

func TestArrayValidator_DifferentArrayTypes(t *testing.T) {
	validator := &ArrayValidator[any]{}

	// Test with different array types
	testCases := []interface{}{
		[]string{"a", "b", "c"},
		[]int{1, 2, 3},
		[]bool{true, false},
		[]float64{1.1, 2.2, 3.3},
		[]interface{}{"mixed", 42, true},
	}

	for _, testCase := range testCases {
		result := validator.Validate(testCase)
		if !result.IsValid {
			t.Errorf("Array validator should accept array of type %T: %v", testCase, testCase)
		}
	}
}

func TestArrayValidator_ArrayType(t *testing.T) {
	validator := &ArrayValidator[any]{}

	// Test with actual array type (not slice)
	array := [3]int{1, 2, 3}
	result := validator.Validate(array)
	if !result.IsValid {
		t.Error("Array validator should accept array type")
	}
}

func TestArrayValidator_ItemValidatorWithNestedFields(t *testing.T) {
	// Create a string validator with constraints
	itemValidator := &StringValidator{}
	itemValidator.MinLength(2).MaxLength(10)

	validator := &ArrayValidator[any]{
		ItemValidator: itemValidator,
	}

	// Test with valid string array
	validArray := []interface{}{"John", "Jane", "Bob"}
	result := validator.Validate(validArray)
	if !result.IsValid {
		t.Error("Array validator should accept valid string array")
	}

	// Test with invalid string array (contains invalid strings)
	invalidArray := []interface{}{"John", "A", "too_long_string_for_validation"}
	result = validator.Validate(invalidArray)
	if result.IsValid {
		t.Error("Array validator should reject array with invalid strings")
	}
	if len(result.Errors) < 2 {
		t.Errorf("Expected at least 2 errors, got %d", len(result.Errors))
	}

	// Check that errors have proper field paths
	for _, err := range result.Errors {
		if err.Field == "" {
			t.Error("Expected field path in error")
		}
		// Should have format like "[1]" or "[2]"
		if len(err.Field) < 3 {
			t.Errorf("Expected proper field path, got '%s'", err.Field)
		}
	}

	// Check that errors are reported for the correct indices
	errorFields := make(map[string]bool)
	for _, err := range result.Errors {
		errorFields[err.Field] = true
	}

	if !errorFields["[1]"] {
		t.Error("Expected error for index [1] (too short string)")
	}
	if !errorFields["[2]"] {
		t.Error("Expected error for index [2] (too long string)")
	}
}
