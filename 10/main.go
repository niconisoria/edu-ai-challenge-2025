package main

import (
	"fmt"
	"os"
)

func main() {
	// Initialize configuration
	cfg := NewConfig()
	if err := cfg.Load(); err != nil {
		fmt.Printf("Error loading configuration: %v\n", err)
		os.Exit(1)
	}

	// Initialize CLI
	app, err := NewCLI(cfg)
	if err != nil {
		fmt.Printf("Error initializing CLI: %v\n", err)
		os.Exit(1)
	}

	// Run the application
	if err := app.Run(); err != nil {
		fmt.Printf("Error running application: %v\n", err)
		os.Exit(1)
	}
}
