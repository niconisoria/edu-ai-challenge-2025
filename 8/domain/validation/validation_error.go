package validation

import (
	"fmt"
)

// ValidationError represents a validation error with field path and message
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("%s: %s", e.Field, e.Message)
	}
	return e.Message
}

// ValidationResult contains validation results
type ValidationResult struct {
	IsValid bool
	Errors  []ValidationError
}
