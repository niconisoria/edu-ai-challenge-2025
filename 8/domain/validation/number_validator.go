package validation

import (
	"fmt"
)

// NumberValidator validates numeric values
type NumberValidator struct {
	BaseValidator
	min *float64
	max *float64
}

func (n *NumberValidator) Validate(value any) ValidationResult {
	// Handle nil values for optional validation
	if value == nil {
		if n.isOptional() {
			return ValidationResult{IsValid: true, Errors: nil}
		}
		return ValidationResult{
			IsValid: false,
			Errors: []ValidationError{
				{
					Field:   "",
					Message: n.getMessage("Number value is required"),
				},
			},
		}
	}

	// Convert value to float64 for validation
	var numValue float64
	switch v := value.(type) {
	case float64:
		numValue = v
	case float32:
		numValue = float64(v)
	case int:
		numValue = float64(v)
	case int8:
		numValue = float64(v)
	case int16:
		numValue = float64(v)
	case int32:
		numValue = float64(v)
	case int64:
		numValue = float64(v)
	case uint:
		numValue = float64(v)
	case uint8:
		numValue = float64(v)
	case uint16:
		numValue = float64(v)
	case uint32:
		numValue = float64(v)
	case uint64:
		numValue = float64(v)
	default:
		return ValidationResult{
			IsValid: false,
			Errors: []ValidationError{
				{
					Field:   "",
					Message: n.getMessage(fmt.Sprintf("Expected numeric value, got %T", value)),
				},
			},
		}
	}

	// Check min constraint
	if n.min != nil && numValue < *n.min {
		return ValidationResult{
			IsValid: false,
			Errors: []ValidationError{
				{
					Field:   "",
					Message: n.getMessage(fmt.Sprintf("Number must be at least %f", *n.min)),
				},
			},
		}
	}

	// Check max constraint
	if n.max != nil && numValue > *n.max {
		return ValidationResult{
			IsValid: false,
			Errors: []ValidationError{
				{
					Field:   "",
					Message: n.getMessage(fmt.Sprintf("Number must be at most %f", *n.max)),
				},
			},
		}
	}

	// Value is valid
	return ValidationResult{IsValid: true, Errors: nil}
}

func (n *NumberValidator) Min(min float64) *NumberValidator {
	n.min = &min
	return n
}

func (n *NumberValidator) Max(max float64) *NumberValidator {
	n.max = &max
	return n
}

func (n *NumberValidator) Optional() Validator[float64] {
	n.setOptional()
	return n
}

func (n *NumberValidator) WithMessage(message string) Validator[float64] {
	n.setMessage(message)
	return n
}
