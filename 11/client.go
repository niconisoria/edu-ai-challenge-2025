package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

// ModelConfig holds configuration for different OpenAI models
// Only Model is kept for simplicity
type ModelConfig struct {
	Model string
}

// Client wraps the OpenAI client with additional configuration
type Client struct {
	client openai.Client
	config ModelConfig
	ctx    context.Context
	cancel context.CancelFunc
}

// DefaultModelConfigs provides predefined configurations for common models
var DefaultModelConfigs = map[string]ModelConfig{
	"gpt-4.1-mini": {
		Model: openai.ChatModelGPT4_1Mini,
	},
	"whisper-1": {
		Model: openai.AudioModelWhisper1,
	},
}

// NewClient creates a new OpenAI client with the specified model configuration
func NewClient(modelName string) (*Client, error) {
	// Check if API key is set
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("OPENAI_API_KEY environment variable is required")
	}

	// Get model configuration
	config, exists := DefaultModelConfigs[modelName]
	if !exists {
		return nil, fmt.Errorf("unknown model: %s", modelName)
	}

	// Create OpenAI client with API key and base URL
	client := openai.NewClient(option.WithAPIKey(apiKey))

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	return &Client{
		client: client,
		config: config,
		ctx:    ctx,
		cancel: cancel,
	}, nil
}

// CreateChatCompletion creates a chat completion with a system and user message (no structured output)
func (c *Client) CreateChatCompletion(systemMessage string, userMessage string) (*openai.ChatCompletion, error) {
	messages := []openai.ChatCompletionMessageParamUnion{
		openai.SystemMessage(systemMessage),
		openai.UserMessage(userMessage),
	}

	req := openai.ChatCompletionNewParams{
		Model:    c.config.Model,
		Messages: messages,
	}

	return c.client.Chat.Completions.New(c.ctx, req)
}

// CreateChatCompletionWithStructuredOutput creates a chat completion with a system and user message and a JSON Schema for structured output
func (c *Client) CreateChatCompletionWithStructuredOutput(systemMessage string, userMessage string, jsonSchema openai.ResponseFormatJSONSchemaJSONSchemaParam) (*openai.ChatCompletion, error) {
	messages := []openai.ChatCompletionMessageParamUnion{
		openai.SystemMessage(systemMessage),
		openai.UserMessage(userMessage),
	}

	responseFormat := openai.ChatCompletionNewParamsResponseFormatUnion{
		OfJSONSchema: &openai.ResponseFormatJSONSchemaParam{
			JSONSchema: jsonSchema,
		},
	}

	req := openai.ChatCompletionNewParams{
		Model:          c.config.Model,
		Messages:       messages,
		ResponseFormat: responseFormat,
	}

	return c.client.Chat.Completions.New(c.ctx, req)
}

// CreateAudioTranscriptionWithOptions creates an audio transcription with additional options
func (c *Client) CreateAudioTranscriptionWithOptions(audioFilePath string, language string, prompt string) (*openai.Transcription, error) {
	// Check if the model is Whisper
	if c.config.Model != "whisper-1" {
		return nil, fmt.Errorf("audio transcription requires whisper-1 model, got: %s", c.config.Model)
	}

	// Open the audio file
	audioFile, err := os.Open(audioFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open audio file: %w", err)
	}
	defer audioFile.Close()

	// Create transcription request with options
	req := openai.AudioTranscriptionNewParams{
		File:     audioFile,
		Model:    openai.AudioModelWhisper1,
		Language: openai.String(language),
		Prompt:   openai.String(prompt),
	}

	transcription, err := c.client.Audio.Transcriptions.New(c.ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create transcription: %w", err)
	}

	return transcription, nil
}

// Close closes the client context
func (c *Client) Close() {
	if c.cancel != nil {
		c.cancel()
	}
}
