package main

import (
	"fmt"

	"github.com/srowles/aoc2024"
)

func main() {
	input := aoc2024.InputFromWebsite("2")
	lines := aoc2024.Lines(input)

	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

func part1(lines []string) int {
	safeCount := 0
	for _, l := range lines {
		ints := aoc2024.Slice(l, " ", aoc2024.Int)
		if safe(ints) {
			safeCount++
		}
	}

	return safeCount
}

func part2(lines []string) int {
	safeCount := 0
outer:
	for _, l := range lines {
		ints := aoc2024.Slice(l, " ", aoc2024.Int)
		if safe(ints) {
			safeCount++
			continue
		}
		for i := 0; i < len(ints); i++ {
			sub := make([]int64, len(ints))
			copy(sub, ints)
			sub = append(sub[:i], sub[i+1:]...)
			if safe(sub) {
				safeCount++
				continue outer
			}
		}
	}

	return safeCount
}

func safe(ints []int64) bool {
	diff := int64(0)
	for i := 1; i < len(ints); i++ {
		if ints[i] == ints[i-1] {
			return false
		}
		ndiff := ints[i] - ints[i-1]
		if diff > 0 && ndiff < 0 {
			return false
		}
		if diff < 0 && ndiff > 0 {
			return false
		}
		diff = ndiff

		if aoc2024.Abs(diff) > 3 {
			return false
		}
	}
	return true
}
