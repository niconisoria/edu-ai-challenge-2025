package validation

import (
	"fmt"
	"reflect"
)

// BooleanValidator validates boolean values
type BooleanValidator struct {
	BaseValidator
}

func (b *BooleanValidator) Validate(value any) ValidationResult {
	// Handle nil values for optional validation
	if value == nil {
		if b.isOptional() {
			return ValidationResult{IsValid: true, Errors: nil}
		}
		return ValidationResult{
			IsValid: false,
			Errors: []ValidationError{
				{
					Field:   "",
					Message: b.getMessage("Boolean value is required"),
				},
			},
		}
	}

	// Check if the value is actually a boolean
	if reflect.TypeOf(value).Kind() != reflect.Bool {
		return ValidationResult{
			IsValid: false,
			Errors: []ValidationError{
				{
					Field:   "",
					Message: b.getMessage(fmt.Sprintf("Expected boolean value, got %T", value)),
				},
			},
		}
	}

	// Value is a valid boolean
	return ValidationResult{IsValid: true, Errors: nil}
}

func (b *BooleanValidator) Optional() Validator[bool] {
	b.setOptional()
	return b
}

func (b *BooleanValidator) WithMessage(message string) Validator[bool] {
	b.setMessage(message)
	return b
}
