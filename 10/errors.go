package main

import (
	"fmt"
)

// AppError represents application-specific errors
type AppError struct {
	Type    string
	Message string
	Err     error
}

// Error types
const (
	ErrorTypeConfig  = "configuration"
	ErrorTypeProduct = "product"
	ErrorTypeAI      = "ai"
)

// Error implementation
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s error: %s: %v", e.Type, e.Message, e.Err)
	}
	return fmt.Sprintf("%s error: %s", e.Type, e.Message)
}

func (e *AppError) Unwrap() error {
	return e.Err
}

// Helper functions to create errors
func NewConfigError(message string, err error) *AppError {
	return &AppError{Type: ErrorTypeConfig, Message: message, Err: err}
}

func NewProductError(message string, err error) *AppError {
	return &AppError{Type: ErrorTypeProduct, Message: message, Err: err}
}

func NewAIError(message string, err error) *AppError {
	return &AppError{Type: ErrorTypeAI, Message: message, Err: err}
}
