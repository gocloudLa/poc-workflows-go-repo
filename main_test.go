package main

import "testing"

func TestGreet(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty name", "", "Hello, World!"},
		{"with name", "Go", "Hello, Go!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := greet(tt.input)
			if result != tt.expected {
				t.Errorf("greet(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative and positive", -1, 1, 0},
		{"large numbers", 50, 50, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("add(%d, %d) = %d, want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"positive numbers", 2, 3, 6},
		{"zero", 0, 5, 0},
		{"same numbers", 5, 5, 25},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("multiply(%d, %d) = %d, want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

