package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func prepareProcessor(processor *Processor) {
	// Load schema from file
	err := processor.LoadSchemaFromFile("schema.json")
	if err != nil {
		fmt.Printf("Warning: Could not load schema from schema.json: %v\n", err)
		fmt.Println("Continuing without schema...")
	}

	// Load system prompt from file
	systemPrompt, err := os.ReadFile("system.txt")
	if err != nil {
		fmt.Printf("Warning: Could not load system prompt from system.txt: %v\n", err)
		fmt.Println("Continuing without system prompt...")
	} else {
		processor.SetSystemPrompt(string(systemPrompt))
	}

	// Load additional context from file
	additionalContext, err := os.ReadFile("context.txt")
	if err != nil {
		fmt.Printf("Warning: Could not load additional context from context.txt: %v\n", err)
		fmt.Println("Continuing without additional context...")
	} else {
		processor.SetAdditionalContext(string(additionalContext))
	}

	// Load examples from file
	err = processor.LoadExamplesFromFile("examples.json")
	if err != nil {
		fmt.Printf("Warning: Could not load examples from JSON file: %v\n", err)
		fmt.Println("Continuing without examples...")
	}
}

func main() {
	// Initialize processor
	processor, err := NewProcessor()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		fmt.Println("Please set your OpenAI API key:")
		fmt.Println("export OPENAI_API_KEY=your_api_key_here")
		os.Exit(1)
	}

	// Prepare the processor with examples
	prepareProcessor(processor)

	fmt.Println("Welcome to the OpenAI CLI Application!")
	fmt.Println("Enter 'quit' to exit")
	fmt.Println("Responses will be written to response.md")
	fmt.Println("----------------------------")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		if input == "quit" || input == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		if input == "" {
			continue
		}

		// Process the input using OpenAI API
		response, err := processor.ProcessInput(input)
		if err != nil {
			fmt.Printf("Error processing input: %v\n", err)
			continue
		}

		// Write response to file
		err = WriteResponseToFile(input, response)
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
		} else {
			fmt.Println("Response written to response.md")
		}

		// Display formatted response to console
		WriteResponseToConsole(input, response)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}
