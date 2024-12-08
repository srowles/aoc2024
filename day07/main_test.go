package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	t.Parallel()
	// Config

	// Test Cases
	tests := map[string]struct {
		input    string
		expected int64
	}{
		"ex1": {
			input:    "190: 10 19",
			expected: 190,
		},
		"ex2": {
			input:    "3267: 81 40 27",
			expected: 3267,
		},
		"ex3": {
			input:    "292: 11 6 16 20",
			expected: 292,
		},
		"ex4": {
			input:    "83: 17 5",
			expected: 0,
		},
		"exall": {
			input: `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`,
			expected: 3749,
		},
	}

	// Testing
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := part1(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
