package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

// AIProviderImpl implements the AIProvider interface
type AIProviderImpl struct {
	client openai.Client
	model  string
}

// NewAIProvider creates a new AI provider
func NewAIProvider(apiKey, model string) AIProvider {
	return &AIProviderImpl{
		client: openai.NewClient(option.WithAPIKey(apiKey)),
		model:  model,
	}
}

// GetFilterFromQuery extracts filter parameters from a natural language query
func (ai *AIProviderImpl) GetFilterFromQuery(ctx context.Context, query string) (*ProductFilter, error) {
	resp, err := ai.client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Model: ai.model,
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(getSystemMessage()),
			openai.UserMessage(query),
		},
		Tools: []openai.ChatCompletionToolParam{
			{Function: getFilterFunctionSchema()},
		},
	})
	if err != nil {
		return nil, NewAIError("failed to get filter from AI", err)
	}

	if len(resp.Choices) == 0 || len(resp.Choices[0].Message.ToolCalls) == 0 {
		return nil, NewAIError("expected tool call but none was received", nil)
	}

	// Parse the function call arguments
	var filter ProductFilter
	toolCall := resp.Choices[0].Message.ToolCalls[0]
	if err := json.Unmarshal([]byte(toolCall.Function.Arguments), &filter); err != nil {
		return nil, NewAIError("failed to parse function arguments", err)
	}

	return &filter, nil
}

// FormatProductsResponse formats a list of products using AI
func (ai *AIProviderImpl) FormatProductsResponse(ctx context.Context, products []Product) (string, error) {
	// Convert products to filtered products format
	filteredProducts := ConvertProductsToFilteredProducts(products)

	resp, err := ai.client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Model: ai.model,
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.AssistantMessage(fmt.Sprintf("List only the relevant products from lowest price to highest price: %v", filteredProducts)),
		},
		ResponseFormat: openai.ChatCompletionNewParamsResponseFormatUnion{
			OfJSONSchema: &openai.ResponseFormatJSONSchemaParam{
				JSONSchema: GetStructuredOutputSchema(),
			},
		},
	})
	if err != nil {
		return "", NewAIError("failed to get formatted response", err)
	}

	if len(resp.Choices) == 0 {
		return "", NewAIError("no response from OpenAI", nil)
	}

	// Parse the structured response
	structuredResponse, err := ParseStructuredResponse(resp.Choices[0].Message.Content)
	if err != nil {
		return "", NewAIError("failed to parse structured response", err)
	}

	// Format the response using the formatter
	formatter := NewResponseFormatter()
	return formatter.FormatProducts(ConvertFilteredProductsToProducts(structuredResponse.Products)), nil
}

// getSystemMessage returns the system message for the AI assistant
func getSystemMessage() string {
	categories := strings.Join(AvailableCategories(), ", ")
	return fmt.Sprintf(`You are a helpful product search assistant. 
	When users ask for products, use the filter_products function to find matching items.
	Available categories: %s.`, categories)
}

// getFilterFunctionSchema returns the function schema for product filtering
func getFilterFunctionSchema() openai.FunctionDefinitionParam {
	return openai.FunctionDefinitionParam{
		Name:        "filter_products",
		Description: openai.String("Selects products based on the user's input criteria"),
		Parameters: openai.FunctionParameters{
			"type": "object",
			"properties": map[string]interface{}{
				"category": map[string]interface{}{
					"type":        "string",
					"description": fmt.Sprintf("The category of the product (available: %s)", strings.Join(AvailableCategories(), ", ")),
				},
				"max_price": map[string]interface{}{
					"type":        "number",
					"description": "The maximum price of the product",
				},
				"min_rating": map[string]interface{}{
					"type":        "number",
					"description": "The minimum rating of the product (1-5)",
				},
				"in_stock": map[string]interface{}{
					"type":        "boolean",
					"description": "Whether the product is in stock",
				},
			},
			"required": []string{"category", "max_price", "min_rating", "in_stock"},
		},
	}
}
