package main

import (
	"context"
)

// ProductServiceImpl implements the ProductService interface
type ProductServiceImpl struct {
	aiProvider        AIProvider
	productRepo       ProductRepository
	responseFormatter ResponseFormatter
}

// NewProductService creates a new product service
func NewProductService(aiProvider AIProvider, productRepo ProductRepository) ProductService {
	return &ProductServiceImpl{
		aiProvider:        aiProvider,
		productRepo:       productRepo,
		responseFormatter: NewResponseFormatter(),
	}
}

// SearchProducts searches for products based on a natural language query
func (ps *ProductServiceImpl) SearchProducts(ctx context.Context, query string) (*ProductSearchResponse, error) {
	// Step 1: Get filter parameters from AI
	filter, err := ps.aiProvider.GetFilterFromQuery(ctx, query)
	if err != nil {
		return nil, err
	}

	// Step 2: Filter products using the repository
	products, err := ps.productRepo.FilterProducts(*filter)
	if err != nil {
		return nil, err
	}

	// Step 3: Format the response using AI
	message, err := ps.aiProvider.FormatProductsResponse(ctx, products)
	if err != nil {
		return nil, err
	}

	return &ProductSearchResponse{
		Products: products,
		Message:  message,
	}, nil
}

// ResponseFormatterImpl implements the ResponseFormatter interface
type ResponseFormatterImpl struct{}

// NewResponseFormatter creates a new response formatter
func NewResponseFormatter() ResponseFormatter {
	return &ResponseFormatterImpl{}
}

// FormatProducts formats a list of products into a readable string
func (rf *ResponseFormatterImpl) FormatProducts(products []Product) string {
	return formatProductList(products)
}
