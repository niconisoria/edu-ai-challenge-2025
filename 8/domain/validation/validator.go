package validation

// Validator is a generic interface for validating values of type T
type Validator[T any] interface {
	Validate(value any) ValidationResult
	Optional() Validator[T]
	WithMessage(message string) Validator[T]
}

// AnyValidator is an interface that can accept any validator type
type AnyValidator interface {
	Validate(value any) ValidationResult
}
