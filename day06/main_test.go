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
		expected int
	}{
		"ex1": {
			input: `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`,
			expected: 41,
		},
	}

	// Testing
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := part1(test.input)
			require.Equal(t, test.expected, len(actual))
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
			input: `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`,
			expected: 6,
		},
	}

	// Testing
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := part2(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
