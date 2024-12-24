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
		blink    int
		expected int64
	}{
		"ex1": {
			input:    `0 1 10 99 999`,
			blink:    1,
			expected: 7,
		},
		"ex2": {
			input:    `125 17`,
			blink:    6,
			expected: 22,
		},
		"ex2.1": {
			input:    `125 17`,
			blink:    25,
			expected: 55312,
		},
	}

	// Testing
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			l := createList(test.input)
			l = blinkn(l, test.blink)
			assert.Equal(t, test.expected, stones(l))
		})
	}
}
