package validation

import (
	"testing"
)

func TestObjectValidator_ValidateEmptySchema(t *testing.T) {
	validator := NewObjectValidator[map[string]any]()

	// Test with valid object when no schema is defined
	testObject := map[string]any{
		"name": "John",
		"age":  30,
		"city": "New York",
	}

	result := validator.Validate(testObject)
	if !result.IsValid {
		t.Error("Object validator should accept any object when no schema is defined")
	}
	if len(result.Errors) > 0 {
		t.Error("Object validator should not return errors when no schema is defined")
	}
}

func TestObjectValidator_ValidateWithSchema(t *testing.T) {
	// Create a schema with validators
	stringValidator := &StringValidator{}
	numberValidator := &NumberValidator{}

	schema := map[string]AnyValidator{
		"name": stringValidator,
		"age":  numberValidator,
	}

	validator := NewObjectValidator[map[string]any]()
	validator.Schema = schema

	// Test with valid object matching schema
	testObject := map[string]any{
		"name": "John",
		"age":  30,
	}

	result := validator.Validate(testObject)
	if !result.IsValid {
		t.Error("Object validator should accept valid object matching schema")
	}
	if len(result.Errors) > 0 {
		t.Error("Object validator should not return errors for valid object")
	}
}

func TestObjectValidator_ValidateMissingRequiredField(t *testing.T) {
	stringValidator := &StringValidator{}
	numberValidator := &NumberValidator{}

	schema := map[string]AnyValidator{
		"name": stringValidator,
		"age":  numberValidator,
	}

	validator := NewObjectValidator[map[string]any]()
	validator.Schema = schema

	// Test with missing required field
	testObject := map[string]any{
		"name": "John",
		// "age" is missing
	}

	result := validator.Validate(testObject)
	if result.IsValid {
		t.Error("Object validator should reject object with missing required field")
	}
	if len(result.Errors) == 0 {
		t.Error("Object validator should return error for missing required field")
	}

	// Check error details
	foundAgeError := false
	for _, err := range result.Errors {
		if err.Field == "age" {
			foundAgeError = true
			break
		}
	}
	if !foundAgeError {
		t.Error("Object validator should return error for missing 'age' field")
	}
}

func TestObjectValidator_ValidateOptionalField(t *testing.T) {
	stringValidator := &StringValidator{}
	numberValidator := &NumberValidator{}
	numberValidator.Optional() // Make age optional

	schema := map[string]AnyValidator{
		"name": stringValidator,
		"age":  numberValidator,
	}

	validator := NewObjectValidator[map[string]any]()
	validator.Schema = schema

	// Test with missing optional field
	testObject := map[string]any{
		"name": "John",
		// "age" is missing but optional
	}

	result := validator.Validate(testObject)
	if !result.IsValid {
		t.Error("Object validator should accept object with missing optional field")
	}
	if len(result.Errors) > 0 {
		t.Error("Object validator should not return errors for missing optional field")
	}
}

func TestObjectValidator_ValidateExtraField(t *testing.T) {
	stringValidator := &StringValidator{}
	numberValidator := &NumberValidator{}

	schema := map[string]AnyValidator{
		"name": stringValidator,
		"age":  numberValidator,
	}

	validator := NewObjectValidator[map[string]any]()
	validator.Schema = schema

	// Test with extra field not in schema
	testObject := map[string]any{
		"name": "John",
		"age":  30,
		"city": "New York", // Extra field
	}

	result := validator.Validate(testObject)
	if result.IsValid {
		t.Error("Object validator should reject object with extra field")
	}
	if len(result.Errors) == 0 {
		t.Error("Object validator should return error for extra field")
	}

	// Check error details
	foundCityError := false
	for _, err := range result.Errors {
		if err.Field == "city" {
			foundCityError = true
			break
		}
	}
	if !foundCityError {
		t.Error("Object validator should return error for extra 'city' field")
	}
}

func TestObjectValidator_ValidateInvalidFieldType(t *testing.T) {
	stringValidator := &StringValidator{}
	numberValidator := &NumberValidator{}

	schema := map[string]AnyValidator{
		"name": stringValidator,
		"age":  numberValidator,
	}

	validator := NewObjectValidator[map[string]any]()
	validator.Schema = schema

	// Test with invalid field type
	testObject := map[string]any{
		"name": "John",
		"age":  "thirty", // Should be number, not string
	}

	result := validator.Validate(testObject)
	if result.IsValid {
		t.Error("Object validator should reject object with invalid field type")
	}
	if len(result.Errors) == 0 {
		t.Error("Object validator should return error for invalid field type")
	}

	// Check error details
	foundAgeError := false
	for _, err := range result.Errors {
		if err.Field == "age" {
			foundAgeError = true
			break
		}
	}
	if !foundAgeError {
		t.Error("Object validator should return error for invalid 'age' field type")
	}
}

func TestObjectValidator_ValidateNil(t *testing.T) {
	validator := NewObjectValidator[map[string]any]()

	// Test with nil object (not optional)
	result := validator.Validate(nil)
	if result.IsValid {
		t.Error("Object validator should reject nil object when not optional")
	}
	if len(result.Errors) == 0 {
		t.Error("Object validator should return error for nil object")
	}
}

func TestObjectValidator_ValidateNilOptional(t *testing.T) {
	validator := NewObjectValidator[map[string]any]()

	// Test with nil object (optional)
	validator.Optional()
	result := validator.Validate(nil)
	if !result.IsValid {
		t.Error("Object validator should accept nil object when optional")
	}
	if len(result.Errors) > 0 {
		t.Error("Object validator should not return errors for nil object when optional")
	}
}

func TestObjectValidator_ValidateInvalidType(t *testing.T) {
	validator := NewObjectValidator[map[string]any]()

	// Test with invalid types
	invalidTestCases := []interface{}{
		"not an object",
		"",
		42,
		3.14,
		true,
		false,
		[]string{"item1", "item2"},
		&ObjectValidator[map[string]any]{},
	}

	for _, testCase := range invalidTestCases {
		result := validator.Validate(testCase)
		if result.IsValid {
			t.Errorf("Object validator should reject invalid type: %v (%T)", testCase, testCase)
		}
		if len(result.Errors) == 0 {
			t.Errorf("Object validator should return error for invalid type: %v (%T)", testCase, testCase)
		}
	}
}

func TestObjectValidator_ValidateEmptyObject(t *testing.T) {
	validator := NewObjectValidator[map[string]any]()

	// Test with empty object
	emptyObject := map[string]any{}

	result := validator.Validate(emptyObject)
	if !result.IsValid {
		t.Error("Object validator should accept empty object when no schema is defined")
	}
}

func TestObjectValidator_ValidateEmptyObjectWithSchema(t *testing.T) {
	stringValidator := &StringValidator{}
	numberValidator := &NumberValidator{}

	schema := map[string]AnyValidator{
		"name": stringValidator,
		"age":  numberValidator,
	}

	validator := NewObjectValidator[map[string]any]()
	validator.Schema = schema

	// Test with empty object when schema requires fields
	emptyObject := map[string]any{}

	result := validator.Validate(emptyObject)
	if result.IsValid {
		t.Error("Object validator should reject empty object when schema requires fields")
	}
	if len(result.Errors) == 0 {
		t.Error("Object validator should return errors for empty object with required fields")
	}

	// Should have errors for both required fields
	if len(result.Errors) != 2 {
		t.Errorf("Expected 2 errors for missing required fields, got %d", len(result.Errors))
	}
}

func TestObjectValidator_ValidateComplexSchema(t *testing.T) {
	// Create a complex schema with nested validation
	stringValidator := &StringValidator{}
	stringValidator.MinLength(2).MaxLength(50)

	numberValidator := &NumberValidator{}
	numberValidator.Min(0).Max(150)

	booleanValidator := &BooleanValidator{}

	schema := map[string]AnyValidator{
		"name":   stringValidator,
		"age":    numberValidator,
		"active": booleanValidator,
	}

	validator := NewObjectValidator[map[string]any]()
	validator.Schema = schema

	// Test with valid complex object
	validObject := map[string]any{
		"name":   "John Doe",
		"age":    30,
		"active": true,
	}

	result := validator.Validate(validObject)
	if !result.IsValid {
		t.Error("Object validator should accept valid complex object")
	}
	if len(result.Errors) > 0 {
		t.Error("Object validator should not return errors for valid complex object")
	}

	// Test with invalid complex object
	invalidObject := map[string]any{
		"name":   "A",   // Too short
		"age":    200,   // Too high
		"active": "yes", // Wrong type
	}

	result = validator.Validate(invalidObject)
	if result.IsValid {
		t.Error("Object validator should reject invalid complex object")
	}
	if len(result.Errors) == 0 {
		t.Error("Object validator should return errors for invalid complex object")
	}

	// Should have multiple errors
	if len(result.Errors) < 3 {
		t.Errorf("Expected at least 3 errors for invalid complex object, got %d", len(result.Errors))
	}
}

func TestObjectValidator_Optional(t *testing.T) {
	validator := NewObjectValidator[map[string]any]()

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

func TestObjectValidator_WithMessage(t *testing.T) {
	validator := NewObjectValidator[map[string]any]()
	customMessage := "Custom object validation message"

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

func TestObjectValidator_Chaining(t *testing.T) {
	validator := NewObjectValidator[map[string]any]()

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

func TestObjectValidator_InterfaceCompliance(t *testing.T) {
	validator := NewObjectValidator[map[string]any]()

	// Test that it implements the Validator interface
	var _ Validator[map[string]any] = validator
}

func TestObjectValidator_DefaultState(t *testing.T) {
	validator := NewObjectValidator[map[string]any]()

	// Test default state
	if validator.isOptional() {
		t.Error("New ObjectValidator should not be optional by default")
	}

	// Test default message behavior
	defaultMsg := "Default message"
	if validator.getMessage(defaultMsg) != defaultMsg {
		t.Error("New ObjectValidator should return default message when no custom message is set")
	}

	// Test that schema is empty by default
	if validator.Schema == nil {
		t.Error("New ObjectValidator should have an empty schema map")
	}
}

func TestObjectValidator_ErrorMessages(t *testing.T) {
	validator := NewObjectValidator[map[string]any]()

	// Test nil value error message
	result := validator.Validate(nil)
	if result.IsValid {
		t.Error("Object validator should reject nil value")
	}
	if len(result.Errors) == 0 {
		t.Error("Object validator should return error for nil value")
	}
	expectedNilMsg := "Object value is required"
	if result.Errors[0].Message != expectedNilMsg {
		t.Errorf("Expected nil error message '%s', got '%s'", expectedNilMsg, result.Errors[0].Message)
	}

	// Test invalid type error message
	result = validator.Validate("string value")
	if result.IsValid {
		t.Error("Object validator should reject string value")
	}
	if len(result.Errors) == 0 {
		t.Error("Object validator should return error for string value")
	}
	expectedTypeMsg := "Expected object value, got string"
	if result.Errors[0].Message != expectedTypeMsg {
		t.Errorf("Expected type error message '%s', got '%s'", expectedTypeMsg, result.Errors[0].Message)
	}
}

func TestObjectValidator_CustomMessage(t *testing.T) {
	validator := NewObjectValidator[map[string]any]()
	customMessage := "Custom object validation message"
	validator.WithMessage(customMessage)

	// Test that custom message is used for invalid type
	result := validator.Validate("not an object")
	if result.IsValid {
		t.Error("Object validator should reject string value")
	}
	if len(result.Errors) == 0 {
		t.Error("Object validator should return error for string value")
	}
	if result.Errors[0].Message != customMessage {
		t.Errorf("Expected custom message '%s', got '%s'", customMessage, result.Errors[0].Message)
	}
}
