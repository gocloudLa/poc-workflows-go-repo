package main

import (
	"fmt"
	"os"
)

// greet returns a greeting message
func greet(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("Hello, %s!", name)
}

// add returns the sum of two integers
func add(a, b int) int {
	return a + b
}

// multiply returns the product of two integers
func multiply(a, b int) int {
	return a * b
}

func main() {
	fmt.Println("Go Service is running")
	fmt.Println("Version: 1.0.0")
	fmt.Println("Health: OK")

	// Keep service running if needed
	if len(os.Args) > 1 && os.Args[1] == "serve" {
		// In a real scenario, this would start an HTTP server
		fmt.Println("Service listening on :8080")
	}
}
