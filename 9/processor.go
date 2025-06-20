package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go/packages/param"
)

// ServiceReport represents the structured output format for service analysis
type ServiceReport struct {
	BriefHistory        string `json:"brief_history"`
	TargetAudience      string `json:"target_audience"`
	CoreFeatures        string `json:"core_features"`
	UniqueSellingPoints string `json:"unique_selling_points"`
	BusinessModel       string `json:"business_model"`
	TechStackInsights   string `json:"tech_stack_insights"`
	PerceivedStrengths  string `json:"perceived_strengths"`
	PerceivedWeaknesses string `json:"perceived_weaknesses"`
}

// Example represents a few-shot learning example
type Example struct {
	Input  string
	Output string
}

// ExamplesData represents the structure of the JSON file
type ExamplesData struct {
	Examples []Example `json:"examples"`
}

// Processor handles OpenAI API interactions with few-shot learning support
type Processor struct {
	client            *openai.Client
	examples          []Example
	systemPrompt      string
	additionalContext string
	schema            map[string]interface{}
}

// NewProcessor returns a new Processor
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
		schema:            map[string]interface{}{},
	}, nil
}

// SetSystemPrompt sets the system prompt for the processor
func (p *Processor) SetSystemPrompt(prompt string) {
	p.systemPrompt = prompt
}

// SetAdditionalContext sets the additional context for the processor
func (p *Processor) SetAdditionalContext(context string) {
	p.additionalContext = context
}

// AddExample adds a few-shot learning example to the processor
func (p *Processor) AddExample(input, output string) {
	p.examples = append(p.examples, Example{
		Input:  input,
		Output: output,
	})
}

// LoadExamplesFromFile loads examples from a file
func (p *Processor) LoadExamplesFromFile(filePath string) error {
	// Read the file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	// Parse the JSON data
	var examplesData ExamplesData
	if err := json.Unmarshal(data, &examplesData); err != nil {
		return fmt.Errorf("failed to parse file %s: %w", filePath, err)
	}

	// Add all examples to the processor
	for _, example := range examplesData.Examples {
		p.AddExample(example.Input, example.Output)
	}

	fmt.Printf("Loaded %d examples from %s\n", len(examplesData.Examples), filePath)
	return nil
}

// LoadSchemaFromFile loads a JSON schema from a file and updates the processor
func (p *Processor) LoadSchemaFromFile(filePath string) error {
	schemaData, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read schema file %s: %w", filePath, err)
	}

	var schema map[string]interface{}
	if err := json.Unmarshal(schemaData, &schema); err != nil {
		return fmt.Errorf("failed to parse schema file %s: %w", filePath, err)
	}

	p.schema = schema
	return nil
}

// ProcessInput processes input using few-shot learning examples and returns structured output
func (p *Processor) ProcessInput(input string) (*ServiceReport, error) {
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
			OfUser: &openai.ChatCompletionUserMessageParam{
				Content: openai.ChatCompletionUserMessageParamContentUnion{
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

	// Prepare completion parameters with structured output
	params := openai.ChatCompletionNewParams{
		Messages:    messages,
		Model:       "gpt-4.1-mini",
		MaxTokens:   openai.Int(2000),
		Temperature: openai.Float(0.7),
		ResponseFormat: openai.ChatCompletionNewParamsResponseFormatUnion{
			OfJSONSchema: &openai.ResponseFormatJSONSchemaParam{
				JSONSchema: openai.ResponseFormatJSONSchemaJSONSchemaParam{
					Name:        "service_report",
					Description: param.Opt[string]{Value: "A structured service analysis report"},
					Schema:      p.schema,
				},
			},
		},
	}

	chatCompletion, err := p.client.Chat.Completions.New(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("OpenAI API error: %w", err)
	}

	if len(chatCompletion.Choices) == 0 {
		return nil, fmt.Errorf("no response from OpenAI")
	}

	// Parse the JSON response into ServiceReport struct
	var serviceReport ServiceReport
	responseContent := chatCompletion.Choices[0].Message.Content
	if err := json.Unmarshal([]byte(responseContent), &serviceReport); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	return &serviceReport, nil
}
