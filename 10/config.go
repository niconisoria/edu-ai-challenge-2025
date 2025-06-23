package main

import (
	"os"
)

// Application constants
const (
	// Default model for OpenAI
	DefaultModel = "gpt-4.1-mini"

	// Available product categories
	CategoryElectronics = "Electronics"
	CategoryFitness     = "Fitness"
	CategoryKitchen     = "Kitchen"
	CategoryBooks       = "Books"
	CategoryClothing    = "Clothing"

	// CLI messages
	WelcomeMessage = "Welcome to the AI Product Search CLI Application!"
	ExitMessage    = "Goodbye!"
	PromptSymbol   = "> "

	// File paths
	ProductsFilePath = "products.json"

	// Environment variables
	OpenAIAPIKeyEnv = "OPENAI_API_KEY"
)

// Config holds the application configuration
type Config struct {
	OpenAIAPIKey string
	Model        string
}

// NewConfig creates a new configuration instance
func NewConfig() *Config {
	return &Config{
		Model: DefaultModel,
	}
}

// Load loads configuration from environment variables
func (c *Config) Load() error {
	apiKey := os.Getenv(OpenAIAPIKeyEnv)
	if apiKey == "" {
		return NewConfigError("OpenAI API key is required", nil)
	}
	c.OpenAIAPIKey = apiKey

	return nil
}

// AvailableCategories returns all available product categories
func AvailableCategories() []string {
	return []string{
		CategoryElectronics,
		CategoryFitness,
		CategoryKitchen,
		CategoryBooks,
		CategoryClothing,
	}
}
