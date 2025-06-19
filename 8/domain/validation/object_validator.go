package validation

import (
	"fmt"
	"reflect"
)

// ObjectValidator validates object values with a schema
type ObjectValidator[T any] struct {
	BaseValidator
	Schema map[string]AnyValidator
}

func (o *ObjectValidator[T]) Validate(value any) ValidationResult {
	// Handle nil values for optional validation
	if value == nil {
		if o.isOptional() {
			return ValidationResult{IsValid: true, Errors: nil}
		}
		return ValidationResult{
			IsValid: false,
			Errors: []ValidationError{
				{
					Field:   "",
					Message: o.getMessage("Object value is required"),
				},
			},
		}
	}

	// Check if the value is actually a map
	if reflect.TypeOf(value).Kind() != reflect.Map {
		return ValidationResult{
			IsValid: false,
			Errors: []ValidationError{
				{
					Field:   "",
					Message: o.getMessage(fmt.Sprintf("Expected object value, got %T", value)),
				},
			},
		}
	}

	// Convert to map[string]any for validation
	objValue, ok := value.(map[string]any)
	if !ok {
		// Try to convert from other map types
		val := reflect.ValueOf(value)
		objValue = make(map[string]any)
		for _, key := range val.MapKeys() {
			objValue[fmt.Sprintf("%v", key.Interface())] = val.MapIndex(key).Interface()
		}
	}

	// If no schema is defined, accept any object
	if len(o.Schema) == 0 {
		return ValidationResult{IsValid: true, Errors: nil}
	}

	// Validate each field in the schema
	var errors []ValidationError
	for fieldName, fieldValidator := range o.Schema {
		fieldValue, exists := objValue[fieldName]

		// If field doesn't exist, check if it's optional
		if !exists {
			// Check if the field validator is optional
			if optionalValidator, ok := fieldValidator.(interface{ isOptional() bool }); ok {
				if optionalValidator.isOptional() {
					continue // Skip optional missing fields
				}
			}

			// Field is required but missing
			errors = append(errors, ValidationError{
				Field:   fieldName,
				Message: o.getMessage(fmt.Sprintf("Field '%s' is required", fieldName)),
			})
			continue
		}

		// Validate the field value
		fieldResult := fieldValidator.Validate(fieldValue)
		if !fieldResult.IsValid {
			// Add field prefix to all errors from this field
			for _, fieldError := range fieldResult.Errors {
				errors = append(errors, ValidationError{
					Field:   fieldName,
					Message: fieldError.Message,
				})
			}
		}
	}

	// Check for extra fields (not in schema)
	for fieldName := range objValue {
		if _, exists := o.Schema[fieldName]; !exists {
			errors = append(errors, ValidationError{
				Field:   fieldName,
				Message: o.getMessage(fmt.Sprintf("Unexpected field '%s'", fieldName)),
			})
		}
	}

	if len(errors) > 0 {
		return ValidationResult{
			IsValid: false,
			Errors:  errors,
		}
	}

	return ValidationResult{IsValid: true, Errors: nil}
}

func (o *ObjectValidator[T]) Optional() Validator[T] {
	o.setOptional()
	return o
}

func (o *ObjectValidator[T]) WithMessage(message string) Validator[T] {
	o.setMessage(message)
	return o
}

// NewObjectValidator creates a new ObjectValidator with initialized schema
func NewObjectValidator[T any]() *ObjectValidator[T] {
	return &ObjectValidator[T]{
		Schema: make(map[string]AnyValidator),
	}
}
