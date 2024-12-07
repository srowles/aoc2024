package main

import (
	"fmt"
	"slices"

	"github.com/srowles/aoc2024"
)

func main() {
	input := aoc2024.InputFromWebsite("5")
	lines := aoc2024.Lines(input)

	fmt.Println(process(lines, false))
	fmt.Println(process(lines, true))
}

func process(lines []string, fix bool) int {
	before := make(map[int64][]int64)
	next := 0
	for i, l := range lines {
		var p, b int64
		_, err := fmt.Sscanf(l, "%d|%d", &p, &b)
		if err != nil {
			next = i
			break
		}

		beforeList, ok := before[p]
		if !ok {
			before[p] = []int64{b}
			continue
		}
		beforeList = append(beforeList, b)
		before[p] = beforeList
	}

	sum := 0
	for i := next + 1; i < len(lines); i++ {
		printInput := aoc2024.Slice(lines[i], ",", aoc2024.Int)
		sortfunc := func(a, b int64) int {
			bl := before[a]
			for _, v := range bl {
				if v == b {
					return -1
				}
			}
			return 0
		}
		printOrdered := slices.IsSortedFunc(printInput, sortfunc)
		if fix {
			if printOrdered {
				continue
			}
			slices.SortFunc(printInput, sortfunc)
		} else {
			if !printOrdered {
				continue
			}
		}
		idx := len(printInput) / 2
		sum += int(printInput[idx])
	}

	return sum
}
