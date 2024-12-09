package main

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	t.Parallel()
	// Config

	// Test Cases
	tests := map[string]struct {
		input    string
		expected int
	}{
		"ex1": {
			input:    `2333133121414131402`,
			expected: 4,
		},
		"ex2": {
			input:    `12345`,
			expected: 4,
		},
	}

	// Testing
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			expanded := expand(test.input)
			fmt.Println(expanded)
			compact(expanded)
			fmt.Println(expanded)
			fmt.Println(sum(expanded))
			t.Fail()
		})
	}
}

func TestPart2(t *testing.T) {
	t.Parallel()
	// Config

	// Test Cases
	tests := map[string]struct {
		input    string
		expected int
	}{
		"ex1": {
			input:    `2333133121414131402`,
			expected: 4,
		},
	}

	// Testing
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			expanded := expand(test.input)
			fmt.Println(expanded)
			compact2(expanded)
			fmt.Println(expanded)
			fmt.Println(sum(expanded))
			t.Fail()
		})
	}
}
