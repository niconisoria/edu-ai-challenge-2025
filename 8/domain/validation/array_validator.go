package validation

import (
	"fmt"
	"reflect"
)

// ArrayValidator validates array values
type ArrayValidator[T any] struct {
	BaseValidator
	ItemValidator AnyValidator
}

func (a *ArrayValidator[T]) Validate(value any) ValidationResult {
	// Handle nil values for optional validation
	if value == nil {
		if a.isOptional() {
			return ValidationResult{IsValid: true, Errors: nil}
		}
		return ValidationResult{
			IsValid: false,
			Errors: []ValidationError{
				{
					Field:   "",
					Message: a.getMessage("Array value is required"),
				},
			},
		}
	}

	// Check if the value is actually a slice/array
	valueType := reflect.TypeOf(value)
	if valueType.Kind() != reflect.Slice && valueType.Kind() != reflect.Array {
		return ValidationResult{
			IsValid: false,
			Errors: []ValidationError{
				{
					Field:   "",
					Message: a.getMessage(fmt.Sprintf("Expected array/slice value, got %T", value)),
				},
			},
		}
	}

	// If no item validator is provided, the array is valid
	if a.ItemValidator == nil {
		return ValidationResult{IsValid: true, Errors: nil}
	}

	// Validate each item in the array
	valueReflect := reflect.ValueOf(value)
	var errors []ValidationError

	for i := 0; i < valueReflect.Len(); i++ {
		item := valueReflect.Index(i).Interface()
		itemResult := a.ItemValidator.Validate(item)

		if !itemResult.IsValid {
			// Add item index to field path for better error reporting
			for _, err := range itemResult.Errors {
				fieldPath := fmt.Sprintf("[%d]", i)
				if err.Field != "" {
					fieldPath = fmt.Sprintf("[%d].%s", i, err.Field)
				}
				errors = append(errors, ValidationError{
					Field:   fieldPath,
					Message: err.Message,
				})
			}
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

func (a *ArrayValidator[T]) Optional() Validator[T] {
	a.setOptional()
	return a
}

func (a *ArrayValidator[T]) WithMessage(message string) Validator[T] {
	a.setMessage(message)
	return a
}
