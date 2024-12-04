package main

import (
	"fmt"

	"github.com/srowles/aoc2024"
)

func main() {
	input := aoc2024.InputFromWebsite("4")
	lines := aoc2024.Lines(input)

	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

var allDirections = []aoc2024.Coord{
	{X: -1, Y: 0},
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: 0, Y: -1},
	{X: 1, Y: 1},
	{X: 1, Y: -1},
	{X: -1, Y: 1},
	{X: -1, Y: -1},
}

var diagonalsOnly = []aoc2024.Coord{
	{X: 1, Y: 1},
	{X: 1, Y: -1},
	{X: -1, Y: 1},
	{X: -1, Y: -1},
}

func part1(lines []string) int {
	grid := make(map[aoc2024.Coord]rune)
	for y, line := range lines {
		for x, r := range line {
			grid[aoc2024.Coord{X: x, Y: y}] = r
		}
	}

	words := loop(grid, "XMAS", allDirections)
	return len(words)
}

func part2(lines []string) int {
	grid := make(map[aoc2024.Coord]rune)
	for y, line := range lines {
		for x, r := range line {
			grid[aoc2024.Coord{X: x, Y: y}] = r
		}
	}

	words := loop(grid, "MAS", diagonalsOnly)
	acount := make(map[aoc2024.Coord]int)
	for _, w := range words {
		acount[w.start.Move(w.direction)] += 1
	}

	count := 0
	for _, c := range acount {
		if c == 2 {
			count++
		}
	}
	return count
}

func loop(grid map[aoc2024.Coord]rune, toFind string, directions []aoc2024.Coord) []word {
	var maxx, maxy int
	for k := range grid {
		if k.X > maxx {
			maxx = k.X
		}
		if k.Y > maxy {
			maxy = k.Y
		}
	}

	words := []word{}
	for y := 0; y <= maxy; y++ {
		for x := 0; x <= maxx; x++ {
			c := aoc2024.Coord{X: x, Y: y}
			if grid[c] == rune(toFind[0]) {
				found := find(grid, toFind[1:], c, directions)
				words = append(words, found...)
			}
		}
	}

	return words
}

type word struct {
	start     aoc2024.Coord
	direction aoc2024.Coord
}

func find(grid map[aoc2024.Coord]rune, toFind string, start aoc2024.Coord, directions []aoc2024.Coord) []word {
	words := []word{}
outer:
	for _, d := range directions {
		loc := start
		for _, l := range toFind {
			loc = loc.Move(d)
			if grid[loc] != l {
				continue outer
			}
		}
		words = append(words, word{start: start, direction: d})
	}

	return words
}
