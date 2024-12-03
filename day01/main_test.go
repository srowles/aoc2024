package main

import (
	"fmt"
	"sort"
	"testing"

	"github.com/srowles/aoc2024"
	"github.com/stretchr/testify/require"
)

func testinput(input string) ([]int, []int) {
	lines := aoc2024.Slice(input, "\n", func(s string) string { return s })
	left := make([]int, len(lines))
	right := make([]int, len(lines))
	for i, line := range lines {
		var l, r int
		fmt.Sscanf(line, "%d   %d", &l, &r)
		left[i] = l
		right[i] = r
	}

	sort.Ints(left)
	sort.Ints(right)

	return left, right
}

func TestPart1(t *testing.T) {
	t.Parallel()
	// Config

	// Test Cases
	tests := map[string]struct {
		input    string
		expected int64
	}{
		"example1": {
			input: `3   4
4   3
2   5
1   3
3   9
3   3`,
			expected: 11,
		},
	}

	// Testing
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := part1(testinput(test.input))
			require.Equal(t, test.expected, actual)
		})
	}
}

func TestPart2(t *testing.T) {
	t.Parallel()
	// Config

	// Test Cases
	tests := map[string]struct {
		input    string
		expected int64
	}{
		"example1": {
			input: `3   4
4   3
2   5
1   3
3   9
3   3`,
			expected: 31,
		},
	}

	// Testing
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := part2(testinput(test.input))
			require.Equal(t, test.expected, actual)
		})
	}
}
