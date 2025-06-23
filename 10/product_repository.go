package main

import (
	"encoding/json"
	"os"
)

// ProductRepositoryImpl implements the ProductRepository interface
type ProductRepositoryImpl struct {
	products []Product
}

// NewProductRepository creates a new product repository
func NewProductRepository() (ProductRepository, error) {
	data, err := os.ReadFile(ProductsFilePath)
	if err != nil {
		return nil, NewProductError("failed to read products file", err)
	}

	var products []Product
	if err := json.Unmarshal(data, &products); err != nil {
		return nil, NewProductError("failed to parse products file", err)
	}

	return &ProductRepositoryImpl{products: products}, nil
}

// GetAllProducts returns all products in the repository
func (pr *ProductRepositoryImpl) GetAllProducts() ([]Product, error) {
	return pr.products, nil
}

// FilterProducts filters products based on the given criteria
func (pr *ProductRepositoryImpl) FilterProducts(filter ProductFilter) ([]Product, error) {
	var filtered []Product

	for _, product := range pr.products {
		if matchesFilter(product, filter) {
			filtered = append(filtered, product)
		}
	}

	return filtered, nil
}
