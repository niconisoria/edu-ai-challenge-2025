package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
)

// CLI handles the command-line interface
type CLI struct {
	productService ProductService
}

// NewCLI creates a new CLI instance
func NewCLI(cfg *Config) (*CLI, error) {
	// Initialize AI provider
	aiProvider := NewAIProvider(cfg.OpenAIAPIKey, cfg.Model)

	// Initialize product repository
	productRepo, err := NewProductRepository()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize product repository: %w", err)
	}

	// Initialize product service
	productService := NewProductService(aiProvider, productRepo)

	return &CLI{
		productService: productService,
	}, nil
}

// Run starts the CLI application
func (c *CLI) Run() error {
	fmt.Println(WelcomeMessage)
	fmt.Println("Type 'quit' or 'exit' to leave the application.")
	fmt.Println(strings.Repeat("-", 50))

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(PromptSymbol)
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		switch strings.ToLower(input) {
		case "quit", "exit":
			fmt.Println(ExitMessage)
			return nil
		default:
			c.processInput(input)
		}
	}

	return scanner.Err()
}

// processInput handles user input and generates responses
func (c *CLI) processInput(input string) {
	ctx := context.Background()

	response, err := c.productService.SearchProducts(ctx, input)
	if err != nil {
		fmt.Printf("Error processing input: %v\n", err)
		fmt.Println("Make sure you have set up your OPENAI_API_KEY environment variable.")
		return
	}

	fmt.Println(response.Message)
}
