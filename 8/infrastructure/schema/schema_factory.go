package schema

import (
	"validation-system/domain/validation"
)

// Schema provides factory methods for creating validators
type Schema struct{}

// String creates a new string validator
func (s *Schema) String() *validation.StringValidator {
	return &validation.StringValidator{}
}

// Number creates a new number validator
func (s *Schema) Number() *validation.NumberValidator {
	return &validation.NumberValidator{}
}

// Boolean creates a new boolean validator
func (s *Schema) Boolean() *validation.BooleanValidator {
	return &validation.BooleanValidator{}
}

// Date creates a new date validator
func (s *Schema) Date() *validation.DateValidator {
	return &validation.DateValidator{}
}

// Object creates a new object validator with the given schema
func (s *Schema) Object(schema map[string]validation.AnyValidator) *validation.ObjectValidator[map[string]any] {
	return &validation.ObjectValidator[map[string]any]{Schema: schema}
}

// Array creates a new array validator with the given item validator
func (s *Schema) Array(itemValidator validation.AnyValidator) *validation.ArrayValidator[any] {
	return &validation.ArrayValidator[any]{ItemValidator: itemValidator}
}
