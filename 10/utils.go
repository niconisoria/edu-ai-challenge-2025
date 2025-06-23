package main

import (
	"fmt"
	"strings"
)

// matchesFilter checks if a product matches the given filter criteria
func matchesFilter(product Product, filter ProductFilter) bool {
	return product.Category == filter.Category &&
		product.Price <= filter.MaxPrice &&
		product.Rating >= filter.MinRating &&
		product.InStock == filter.InStock
}

// getStockStatus returns a human-readable stock status
func getStockStatus(inStock bool) string {
	if inStock {
		return "In Stock"
	}
	return "Out of Stock"
}

// formatProductList formats a list of products into a readable string
func formatProductList(products []Product) string {
	if len(products) == 0 {
		return "No products found matching your criteria."
	}

	var builder strings.Builder
	builder.WriteString("Filtered Products:\n")

	for i, product := range products {
		stockStatus := getStockStatus(product.InStock)
		builder.WriteString(fmt.Sprintf("%d. %s - $%.2f, Rating: %.1f, %s\n",
			i+1, product.Name, product.Price, product.Rating, stockStatus))
	}

	return builder.String()
}
