package main

import (
	"fmt"

	"github.com/srowles/aoc2024"
)

func main() {
	input := aoc2024.InputFromWebsite("10")

	book := parseBook(input)
	fmt.Println(scores(book))
	fmt.Println(rate(book))
}

type book struct {
	reachableHeads map[aoc2024.Coord]int
	maxx, maxy     int
	grid           map[aoc2024.Coord]int
}

func (b *book) print() {
	for y := 0; y <= b.maxy; y++ {
		for x := 0; x <= b.maxx; x++ {
			c := aoc2024.Coord{X: x, Y: y}
			if b.grid[c] == -1 {
				fmt.Print(".")
			} else {
				fmt.Print(b.grid[c])
			}
		}
		fmt.Println()
	}
}

var directions = []aoc2024.Coord{
	{X: 0, Y: 1},
	{X: 0, Y: -1},
	{X: -1, Y: 0},
	{X: 1, Y: 0},
}

func (b *book) score(start aoc2024.Coord) {
	if b.grid[start] == 9 {
		b.reachableHeads[start] += 1
		return
	}
	for _, d := range directions {
		next := start.Move(d)
		if b.grid[next] == b.grid[start]+1 {
			b.score(next)
		}
	}
}

func parseBook(input string) *book {
	b := book{
		grid:           make(map[aoc2024.Coord]int),
		reachableHeads: make(map[aoc2024.Coord]int),
	}

	for y, line := range aoc2024.Lines(input) {
		if y > b.maxy {
			b.maxy = y
		}
		for x, v := range line {
			if x > b.maxx {
				b.maxx = x
			}
			if v == '.' {
				b.grid[aoc2024.Coord{X: x, Y: y}] = -1
				continue
			}
			b.grid[aoc2024.Coord{X: x, Y: y}] = int(aoc2024.Int(string(v)))
		}
	}

	return &b
}

func scores(b *book) int {
	score := 0
	for c, v := range b.grid {
		if v == 0 {
			b.reachableHeads = make(map[aoc2024.Coord]int)
			b.score(c)
			score += len(b.reachableHeads)
		}
	}
	return score
}

func rate(b *book) int {
	rating := 0
	for c, v := range b.grid {
		if v == 0 {
			b.reachableHeads = make(map[aoc2024.Coord]int)
			b.score(c)
			for _, v := range b.reachableHeads {
				rating += v
			}
		}
	}
	return rating
}
