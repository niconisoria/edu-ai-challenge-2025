package schema

import (
	"testing"
	"time"
	"validation-system/domain/validation"
)

func TestSchema_String(t *testing.T) {
	schema := &Schema{}

	validator := schema.String()

	if validator == nil {
		t.Error("String() should return a non-nil StringValidator")
	}

	// Test that it's a valid validator by calling a method
	validator.MinLength(1) // This should not panic if it's a proper validator
}

func TestSchema_Number(t *testing.T) {
	schema := &Schema{}

	validator := schema.Number()

	if validator == nil {
		t.Error("Number() should return a non-nil NumberValidator")
	}

	// Test that it's a valid validator by calling a method
	validator.Min(1) // This should not panic if it's a proper validator
}

func TestSchema_Boolean(t *testing.T) {
	schema := &Schema{}

	validator := schema.Boolean()

	if validator == nil {
		t.Error("Boolean() should return a non-nil BooleanValidator")
	}

	// Test that it's a valid validator by calling a method
	result := validator.Validate(true) // This should not panic if it's a proper validator
	if result.IsValid != true {
		t.Error("Boolean validator should validate true as valid")
	}
}

func TestSchema_Date(t *testing.T) {
	schema := &Schema{}

	validator := schema.Date()

	if validator == nil {
		t.Error("Date() should return a non-nil DateValidator")
	}

	// Test that it's a valid validator by calling a method
	result := validator.Validate(time.Now()) // This should not panic if it's a proper validator
	if result.IsValid != true {
		t.Error("Date validator should validate valid date as valid")
	}
}

func TestSchema_Object(t *testing.T) {
	schema := &Schema{}

	// Create a simple object schema
	objectSchema := map[string]validation.AnyValidator{
		"name": schema.String(),
		"age":  schema.Number(),
	}

	validator := schema.Object(objectSchema)

	if validator == nil {
		t.Error("Object() should return a non-nil ObjectValidator")
	}

	// Test that the schema was set correctly
	if len(validator.Schema) != 2 {
		t.Error("Object validator should have 2 schema fields")
	}

	// Test that the validators are of the correct types
	if validator.Schema["name"] == nil {
		t.Error("Name field should be a validator")
	}

	if validator.Schema["age"] == nil {
		t.Error("Age field should be a validator")
	}
}

func TestSchema_Array(t *testing.T) {
	schema := &Schema{}

	// Create an array validator with string items
	itemValidator := schema.String()
	validator := schema.Array(itemValidator)

	if validator == nil {
		t.Error("Array() should return a non-nil ArrayValidator")
	}

	// Test that the item validator was set correctly
	if validator.ItemValidator != itemValidator {
		t.Error("Array validator should have the correct item validator")
	}
}

func TestSchema_ComplexSchema(t *testing.T) {
	schema := &Schema{}

	// Create a complex nested schema
	addressSchema := map[string]validation.AnyValidator{
		"street":     schema.String(),
		"city":       schema.String(),
		"postalCode": schema.String(),
		"country":    schema.String(),
	}

	userSchema := map[string]validation.AnyValidator{
		"id":       schema.String(),
		"name":     schema.String(),
		"email":    schema.String(),
		"age":      schema.Number().Optional(),
		"isActive": schema.Boolean(),
		"tags":     schema.Array(schema.String()),
		"address":  schema.Object(addressSchema).Optional(),
	}

	validator := schema.Object(userSchema)

	if validator == nil {
		t.Error("Complex Object() should return a non-nil ObjectValidator")
	}

	// Test that the schema has the expected number of fields
	if len(validator.Schema) != 7 {
		t.Errorf("Expected 7 fields in user schema, got %d", len(validator.Schema))
	}

	// Test specific field types
	if validator.Schema["id"] == nil {
		t.Error("ID field should be a validator")
	}

	if validator.Schema["age"] == nil {
		t.Error("Age field should be a validator")
	}

	if validator.Schema["isActive"] == nil {
		t.Error("IsActive field should be a validator")
	}

	if validator.Schema["tags"] == nil {
		t.Error("Tags field should be a validator")
	}

	if validator.Schema["address"] == nil {
		t.Error("Address field should be a validator")
	}
}

func TestSchema_EmptyObject(t *testing.T) {
	schema := &Schema{}

	// Create an empty object schema
	emptySchema := map[string]validation.AnyValidator{}

	validator := schema.Object(emptySchema)

	if validator == nil {
		t.Error("Empty Object() should return a non-nil ObjectValidator")
	}

	// Test that the schema is empty
	if len(validator.Schema) != 0 {
		t.Error("Empty object schema should have no fields")
	}
}

func TestSchema_MethodChaining(t *testing.T) {
	schema := &Schema{}

	// Test method chaining on validators
	stringValidator := schema.String().
		MinLength(3).
		MaxLength(50).
		Pattern(`^[a-zA-Z]+$`).
		Optional().
		WithMessage("Custom string message")

	if stringValidator == nil {
		t.Error("Method chaining should return a non-nil validator")
	}

	// Test that it implements the Validator interface
	var _ validation.Validator[string] = stringValidator
}
