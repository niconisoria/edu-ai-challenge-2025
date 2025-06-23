package main

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/openai/openai-go"
)

// Product represents a product in the catalog
type Product struct {
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Rating   float64 `json:"rating"`
	InStock  bool    `json:"in_stock"`
}

// ProductFilter represents the criteria for filtering products
type ProductFilter struct {
	Category  string  `json:"category"`
	MaxPrice  float64 `json:"max_price"`
	MinRating float64 `json:"min_rating"`
	InStock   bool    `json:"in_stock"`
}

// ProductSearchResponse represents the response from product search
type ProductSearchResponse struct {
	Products []Product `json:"products"`
	Message  string    `json:"message"`
}

// FilteredProduct represents a product in the structured output response
type FilteredProduct struct {
	Name    string  `json:"name" jsonschema_description:"The name of the product"`
	Price   float64 `json:"price" jsonschema_description:"The price of the product in dollars"`
	Rating  float64 `json:"rating" jsonschema_description:"The rating of the product (1-5 scale)"`
	InStock bool    `json:"in_stock" jsonschema_description:"Whether the product is currently in stock"`
}

// FilteredProductsResponse represents the structured output response
type FilteredProductsResponse struct {
	Products []FilteredProduct `json:"products" jsonschema_description:"List of filtered products matching the criteria"`
}

// Service interfaces
type ProductService interface {
	SearchProducts(ctx context.Context, query string) (*ProductSearchResponse, error)
}

type AIProvider interface {
	GetFilterFromQuery(ctx context.Context, query string) (*ProductFilter, error)
	FormatProductsResponse(ctx context.Context, products []Product) (string, error)
}

type ProductRepository interface {
	GetAllProducts() ([]Product, error)
	FilterProducts(filter ProductFilter) ([]Product, error)
}

type ResponseFormatter interface {
	FormatProducts(products []Product) string
}

// Conversion functions
func ConvertProductsToFilteredProducts(products []Product) []FilteredProduct {
	var filteredProducts []FilteredProduct
	for _, product := range products {
		filteredProducts = append(filteredProducts, FilteredProduct{
			Name:    product.Name,
			Price:   product.Price,
			Rating:  product.Rating,
			InStock: product.InStock,
		})
	}
	return filteredProducts
}

func ConvertFilteredProductsToProducts(filteredProducts []FilteredProduct) []Product {
	var products []Product
	for _, fp := range filteredProducts {
		products = append(products, Product{
			Name:    fp.Name,
			Price:   fp.Price,
			Rating:  fp.Rating,
			InStock: fp.InStock,
		})
	}
	return products
}

// Schema functions
func GetStructuredOutputSchema() openai.ResponseFormatJSONSchemaJSONSchemaParam {
	reflector := jsonschema.Reflector{
		AllowAdditionalProperties: false,
		DoNotReference:            true,
	}
	var v FilteredProductsResponse
	schema := reflector.Reflect(v)

	return openai.ResponseFormatJSONSchemaJSONSchemaParam{
		Name:        "filtered_products",
		Description: openai.String("Filtered products matching the user's criteria"),
		Schema:      schema,
		Strict:      openai.Bool(true),
	}
}

func ParseStructuredResponse(content string) (*FilteredProductsResponse, error) {
	var response FilteredProductsResponse
	if err := json.Unmarshal([]byte(content), &response); err != nil {
		return nil, err
	}
	return &response, nil
}
