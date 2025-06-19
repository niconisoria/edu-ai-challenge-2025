package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

// Example represents a few-shot learning example
type Example struct {
	Input  string
	Output string
}

// Processor handles OpenAI API interactions with few-shot learning support
type Processor struct {
	client            *openai.Client
	examples          []Example
	systemPrompt      string
	additionalContext string
}

func NewProcessor() (*Processor, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("OPENAI_API_KEY environment variable is required")
	}

	client := openai.NewClient(option.WithAPIKey(apiKey))
	return &Processor{
		client:            &client,
		examples:          make([]Example, 0),
		systemPrompt:      "",
		additionalContext: "",
	}, nil
}

// AddExample adds a few-shot learning example to the processor
func (p *Processor) AddExample(input, output string) {
	p.examples = append(p.examples, Example{
		Input:  input,
		Output: output,
	})
}

// ExamplesData represents the structure of the JSON file
type ExamplesData struct {
	Examples []Example `json:"examples"`
}

// LoadExamplesFromJSON loads examples from a JSON file
func (p *Processor) LoadExamplesFromJSON(filePath string) error {
	// Read the JSON file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read JSON file %s: %w", filePath, err)
	}

	// Parse the JSON data
	var examplesData ExamplesData
	if err := json.Unmarshal(data, &examplesData); err != nil {
		return fmt.Errorf("failed to parse JSON file %s: %w", filePath, err)
	}

	// Add all examples to the processor
	for _, example := range examplesData.Examples {
		p.AddExample(example.Input, example.Output)
	}

	fmt.Printf("Loaded %d examples from %s\n", len(examplesData.Examples), filePath)
	return nil
}

// SetSystemPrompt sets the system prompt for the processor
func (p *Processor) SetSystemPrompt(prompt string) {
	p.systemPrompt = prompt
}

// GetSystemPrompt returns the current system prompt
func (p *Processor) GetSystemPrompt() string {
	return p.systemPrompt
}

// SetAdditionalContext sets the additional context for the processor
func (p *Processor) SetAdditionalContext(context string) {
	p.additionalContext = context
}

// GetAdditionalContext returns the current additional context
func (p *Processor) GetAdditionalContext() string {
	return p.additionalContext
}

// ProcessInput processes input using few-shot learning examples
func (p *Processor) ProcessInput(input string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Build messages array starting with system message
	messages := []openai.ChatCompletionMessageParamUnion{}

	// Add system prompt if set
	if p.systemPrompt != "" {
		messages = append(messages, openai.ChatCompletionMessageParamUnion{
			OfSystem: &openai.ChatCompletionSystemMessageParam{
				Content: openai.ChatCompletionSystemMessageParamContentUnion{
					OfString: openai.String(p.systemPrompt),
				},
			},
		})
	}

	// Add context as a user message if set
	if p.additionalContext != "" {
		messages = append(messages, openai.ChatCompletionMessageParamUnion{
			OfUser: &openai.ChatCompletionUserMessageParam{
				Content: openai.ChatCompletionUserMessageParamContentUnion{
					OfString: openai.String("Context: " + p.additionalContext),
				},
			},
		})
	}

	// Add few-shot examples as user/assistant message pairs
	for _, example := range p.examples {
		// Add user message (input)
		messages = append(messages, openai.ChatCompletionMessageParamUnion{
			OfSystem: &openai.ChatCompletionSystemMessageParam{
				Content: openai.ChatCompletionSystemMessageParamContentUnion{
					OfString: openai.String(example.Input),
				},
			},
		})

		// Add assistant message (output)
		messages = append(messages, openai.ChatCompletionMessageParamUnion{
			OfAssistant: &openai.ChatCompletionAssistantMessageParam{
				Content: openai.ChatCompletionAssistantMessageParamContentUnion{
					OfString: openai.String(example.Output),
				},
			},
		})
	}

	// Add the current user input
	messages = append(messages, openai.ChatCompletionMessageParamUnion{
		OfUser: &openai.ChatCompletionUserMessageParam{
			Content: openai.ChatCompletionUserMessageParamContentUnion{
				OfString: openai.String(input),
			},
		},
	})

	// Prepare completion parameters
	params := openai.ChatCompletionNewParams{
		Messages:    messages,
		Model:       "gpt-4.1-mini",
		MaxTokens:   openai.Int(1000),
		Temperature: openai.Float(0.7),
		ResponseFormat: openai.ChatCompletionNewParamsResponseFormatUnion{
			OfText: &openai.ResponseFormatTextParam{},
		},
	}

	chatCompletion, err := p.client.Chat.Completions.New(ctx, params)
	if err != nil {
		return "", fmt.Errorf("OpenAI API error: %w", err)
	}

	if len(chatCompletion.Choices) == 0 {
		return "", fmt.Errorf("no response from OpenAI")
	}

	return chatCompletion.Choices[0].Message.Content, nil
}

func WriteResponseToFile(input, response string) error {
	// Open file in append mode, create if doesn't exist
	file, err := os.OpenFile("response.md", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write to file
	_, err = file.WriteString(response)
	return err
}
