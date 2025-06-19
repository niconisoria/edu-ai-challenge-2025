package validation

import (
	"testing"
)

func TestBaseValidator_SetOptional(t *testing.T) {
	validator := &BaseValidator{}

	// Initially should not be optional
	if validator.isOptional() {
		t.Error("BaseValidator should not be optional by default")
	}

	// Set as optional
	validator.setOptional()

	if !validator.isOptional() {
		t.Error("BaseValidator should be optional after setOptional()")
	}
}

func TestBaseValidator_SetMessage(t *testing.T) {
	validator := &BaseValidator{}
	customMessage := "Custom validation message"

	// Initially should return default message
	defaultMsg := "Default message"
	result := validator.getMessage(defaultMsg)
	if result != defaultMsg {
		t.Errorf("Expected default message '%s', got '%s'", defaultMsg, result)
	}

	// Set custom message
	validator.setMessage(customMessage)

	result = validator.getMessage(defaultMsg)
	if result != customMessage {
		t.Errorf("Expected custom message '%s', got '%s'", customMessage, result)
	}
}

func TestBaseValidator_GetMessageWithEmptyCustomMessage(t *testing.T) {
	validator := &BaseValidator{}
	defaultMsg := "Default message"

	// Set empty custom message
	validator.setMessage("")

	result := validator.getMessage(defaultMsg)
	if result != defaultMsg {
		t.Errorf("Expected default message '%s' when custom message is empty, got '%s'", defaultMsg, result)
	}
}

func TestBaseValidator_IsOptional(t *testing.T) {
	validator := &BaseValidator{}

	// Test initial state
	if validator.isOptional() {
		t.Error("New BaseValidator should not be optional")
	}

	// Test after setting optional
	validator.setOptional()
	if !validator.isOptional() {
		t.Error("BaseValidator should be optional after setOptional()")
	}
}
