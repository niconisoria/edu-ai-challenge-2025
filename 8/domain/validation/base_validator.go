package validation

// BaseValidator provides common functionality for all validators
type BaseValidator struct {
	optional bool
	message  string
}

func (b *BaseValidator) setOptional() {
	b.optional = true
}

func (b *BaseValidator) setMessage(message string) {
	b.message = message
}

func (b *BaseValidator) getMessage(defaultMsg string) string {
	if b.message != "" {
		return b.message
	}
	return defaultMsg
}

func (b *BaseValidator) isOptional() bool {
	return b.optional
}
