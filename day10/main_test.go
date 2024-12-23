package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
			input: `...0...
		...1...
		...2...
		6543456
		7.....7
		8.....8
		9.....9`,
			expected: 2,
		},
		"ex2": {
			input: `..90..9
		...1.98
		...2..7
		6543456
		765.987
		876....
		987....`,
			expected: 4,
		},
		"exlarger": {
			input: `89010123
		78121874
		87430965
		96549874
		45678903
		32019012
		01329801
		10456732`,
			expected: 36,
		},
	}

	// Testing
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			b := parseBook(test.input)
			// b.print()
			assert.Equal(t, test.expected, scores(b))
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
			input: `.....0.
..4321.
..5..2.
..6543.
..7..4.
..8765.
..9....`,
			expected: 3,
		},
		"ex2": {
			input: `..90..9
...1.98
...2..7
6543456
765.987
876....
987....`,
			expected: 13,
		},
		"ex3": {
			input: `012345
123456
234567
345678
4.6789
56789.`,
			expected: 227,
		},
	}

	// Testing
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			b := parseBook(test.input)
			// b.print()
			assert.Equal(t, test.expected, rate(b))
		})
	}
}
