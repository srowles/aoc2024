package main

import (
	"fmt"
	"strconv"

	"github.com/srowles/aoc2024"
)

func main() {
	input := aoc2024.InputFromWebsite("11")

	l := createList(input)
	l = blinkn(l, 25)
	fmt.Println(stones(l))
	l = blinkn(l, 50)
	fmt.Println(stones(l))
}

func blinkn(l map[int64]int64, n int) map[int64]int64 {
	for i := 0; i < n; i++ {
		l = blink(l)
	}
	return l
}

func blink(l map[int64]int64) map[int64]int64 {
	next := make(map[int64]int64, len(l))
	for k, count := range l {
		if k == 0 {
			next[1] += count
		} else if strval := fmt.Sprintf("%d", k); len(strval)%2 == 0 {
			strval := fmt.Sprintf("%d", k)
			mid := len(strval) / 2
			left, _ := strconv.ParseInt(strval[0:mid], 10, 64)
			right, _ := strconv.ParseInt(strval[mid:], 10, 64)
			next[left] += count
			next[right] += count
		} else {
			n := k * 2024
			next[n] += count
		}
	}

	return next
}

func stones(l map[int64]int64) int64 {
	stones := int64(0)
	for _, v := range l {
		stones += v
	}
	return stones
}

func createList(input string) map[int64]int64 {
	l := make(map[int64]int64)
	for _, v := range aoc2024.Slice(input, " ", aoc2024.Int) {
		l[v] += 1
	}

	return l
}
