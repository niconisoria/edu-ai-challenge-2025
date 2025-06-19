package main

import (
	"fmt"
	"validation-system/domain/validation"
	"validation-system/infrastructure/schema"
)

func main() {
	fmt.Println("Schema Validation Example")
	fmt.Println("=========================")

	ExampleComplexSchema()
}

// ExampleComplexSchema demonstrates complex schema building
func ExampleComplexSchema() {
	// Create schema instance
	schemaBuilder := &schema.Schema{}

	// Define address schema
	addressSchema := schemaBuilder.Object(map[string]validation.AnyValidator{
		"street":     schemaBuilder.String(),
		"city":       schemaBuilder.String(),
		"postalCode": schemaBuilder.String().Pattern(`^\d{5}$`).WithMessage("Postal code must be 5 digits"),
		"country":    schemaBuilder.String(),
	})

	// Define user schema
	userSchema := schemaBuilder.Object(map[string]validation.AnyValidator{
		"id":       schemaBuilder.String().WithMessage("ID must be a string"),
		"name":     schemaBuilder.String().MinLength(2).MaxLength(50),
		"email":    schemaBuilder.String().Pattern(`^[^\s@]+@[^\s@]+\.[^\s@]+$`),
		"age":      schemaBuilder.Number().Optional(),
		"isActive": schemaBuilder.Boolean(),
		"tags":     schemaBuilder.Array(schemaBuilder.String()),
		"address":  addressSchema.Optional(),
		"metadata": schemaBuilder.Object(map[string]validation.AnyValidator{}).Optional(),
	})

	// Test data
	userData := map[string]interface{}{
		"id":       "12345",
		"name":     "John Doe",
		"email":    "john@example.com",
		"isActive": true,
		"tags":     []interface{}{"developer", "designer"},
		"address": map[string]interface{}{
			"street":     "123 Main St",
			"city":       "Anytown",
			"postalCode": "12345",
			"country":    "USA",
		},
	}

	// Validate the data
	result := userSchema.Validate(userData)

	if result.IsValid {
		fmt.Println("User data is valid!")
	} else {
		fmt.Println("User data is invalid:")
		for _, err := range result.Errors {
			fmt.Printf("  - %s: %s\n", err.Field, err.Message)
		}
	}
}
