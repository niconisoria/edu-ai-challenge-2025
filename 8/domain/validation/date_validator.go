package validation

import (
	"fmt"
	"reflect"
	"time"
)

// DateValidator validates date/time values
type DateValidator struct {
	BaseValidator
}

func (d *DateValidator) Validate(value any) ValidationResult {
	// Handle nil values for optional validation
	if value == nil {
		if d.isOptional() {
			return ValidationResult{IsValid: true, Errors: nil}
		}
		return ValidationResult{
			IsValid: false,
			Errors: []ValidationError{
				{
					Field:   "",
					Message: d.getMessage("Date value is required"),
				},
			},
		}
	}

	// Check if the value is actually a time.Time
	if reflect.TypeOf(value) != reflect.TypeOf(time.Time{}) {
		return ValidationResult{
			IsValid: false,
			Errors: []ValidationError{
				{
					Field:   "",
					Message: d.getMessage(fmt.Sprintf("Expected time.Time value, got %T", value)),
				},
			},
		}
	}

	// Value is a valid time.Time
	return ValidationResult{IsValid: true, Errors: nil}
}

func (d *DateValidator) Optional() Validator[time.Time] {
	d.setOptional()
	return d
}

func (d *DateValidator) WithMessage(message string) Validator[time.Time] {
	d.setMessage(message)
	return d
}
