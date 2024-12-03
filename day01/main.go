package main

import (
	"fmt"
	"sort"

	"github.com/srowles/aoc2024"
)

func main() {
	input := aoc2024.InputFromWebsite("1")
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

	fmt.Println(part1(left, right))
	fmt.Println(part2(left, right))
}

func part1(left, right []int) int64 {
	distance := int64(0)
	for i, l := range left {
		distance += int64(aoc2024.Abs(right[i] - l))
	}
	return distance
}

func part2(left, right []int) int64 {
	similarity := int64(0)
	for _, l := range left {
		count := count(right, l)
		similarity += int64(count * l)
	}
	return similarity
}

func count(list []int, val int) int {
	count := 0
	for _, v := range list {
		if v == val {
			count++
		}
		if v > val {
			break
		}
	}
	return count
}
