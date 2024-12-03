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
		input []string
		safe  int
	}{
		"ex1": {
			input: []string{
				"7 6 4 2 1",
				"1 2 7 8 9",
				"9 7 6 2 1",
				"1 3 2 4 5",
				"8 6 4 4 1",
				"1 3 6 7 9",
			},
			safe: 2,
		},
	}

	// Testing
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			safe := part1(test.input)
			require.Equal(t, test.safe, safe)
		})
	}
}

func TestPart2(t *testing.T) {
	t.Parallel()
	// Config

	// Test Cases
	tests := map[string]struct {
		input []string
		safe  int
	}{
		"ex1": {
			input: []string{
				"7 6 4 2 1",
				"1 2 7 8 9",
				"9 7 6 2 1",
				"1 3 2 4 5",
				"8 6 4 4 1",
				"1 3 6 7 9",
			},
			safe: 4,
		},
	}

	// Testing
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			safe := part2(test.input)
			require.Equal(t, test.safe, safe)
		})
	}
}
