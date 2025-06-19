package validation

import (
	"fmt"
	"reflect"
	"regexp"
)

// StringValidator validates string values
type StringValidator struct {
	BaseValidator
	minLength *int
	maxLength *int
	pattern   *regexp.Regexp
}

func (s *StringValidator) Validate(value any) ValidationResult {
	// Handle nil values for optional validation
	if value == nil {
		if s.isOptional() {
			return ValidationResult{IsValid: true, Errors: nil}
		}
		return ValidationResult{
			IsValid: false,
			Errors: []ValidationError{
				{
					Field:   "",
					Message: s.getMessage("String value is required"),
				},
			},
		}
	}

	// Check if the value is actually a string
	if reflect.TypeOf(value).Kind() != reflect.String {
		return ValidationResult{
			IsValid: false,
			Errors: []ValidationError{
				{
					Field:   "",
					Message: s.getMessage(fmt.Sprintf("Expected string value, got %T", value)),
				},
			},
		}
	}

	strValue := value.(string)

	// Check min length constraint
	if s.minLength != nil && len(strValue) < *s.minLength {
		return ValidationResult{
			IsValid: false,
			Errors: []ValidationError{
				{
					Field:   "",
					Message: s.getMessage(fmt.Sprintf("String must be at least %d characters long", *s.minLength)),
				},
			},
		}
	}

	// Check max length constraint
	if s.maxLength != nil && len(strValue) > *s.maxLength {
		return ValidationResult{
			IsValid: false,
			Errors: []ValidationError{
				{
					Field:   "",
					Message: s.getMessage(fmt.Sprintf("String must be at most %d characters long", *s.maxLength)),
				},
			},
		}
	}

	// Check pattern constraint
	if s.pattern != nil && !s.pattern.MatchString(strValue) {
		return ValidationResult{
			IsValid: false,
			Errors: []ValidationError{
				{
					Field:   "",
					Message: s.getMessage(fmt.Sprintf("String must match pattern: %s", s.pattern.String())),
				},
			},
		}
	}

	// Value is valid
	return ValidationResult{IsValid: true, Errors: nil}
}

func (s *StringValidator) MinLength(length int) *StringValidator {
	s.minLength = &length
	return s
}

func (s *StringValidator) MaxLength(length int) *StringValidator {
	s.maxLength = &length
	return s
}

func (s *StringValidator) Pattern(pattern string) *StringValidator {
	regex := regexp.MustCompile(pattern)
	s.pattern = regex
	return s
}

func (s *StringValidator) Optional() Validator[string] {
	s.setOptional()
	return s
}

func (s *StringValidator) WithMessage(message string) Validator[string] {
	s.setMessage(message)
	return s
}
