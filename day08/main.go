package main

import (
	"fmt"

	"github.com/srowles/aoc2024"
)

func main() {
	input := aoc2024.InputFromWebsite("8")

	fmt.Println(nodes(input, false))
	fmt.Println(nodes(input, true))
}

func nodes(input string, repeat bool) int {
	knownFrequencies := make(map[rune][]aoc2024.Coord)
	grid := make(map[aoc2024.Coord]rune)
	var maxx, maxy int
	for y, line := range aoc2024.Lines(input) {
		if y > maxy {
			maxy = y
		}
		for x, r := range line {
			if x > maxx {
				maxx = x
			}
			if r == '.' || r == '#' {
				continue
			}
			grid[aoc2024.Coord{X: x, Y: y}] = r
			knownFrequencies[r] = append(knownFrequencies[r], aoc2024.Coord{X: x, Y: y})
		}
	}
	fmt.Println(maxx, maxy)

	answer := make(map[aoc2024.Coord]rune)
	for k, v := range grid {
		answer[k] = v
	}
	for _, coords := range knownFrequencies {
		if len(coords) == 1 {
			continue
		}
		antigrid := make(map[aoc2024.Coord]bool)
		fmt.Println(coords)
		for i, c := range coords {
			for ii, co := range coords {
				if i == ii {
					continue
				}
				newnodes := antinodes(c, co, repeat)
				for _, n := range newnodes {
					if n.X >= 0 && n.X <= maxx && n.Y >= 0 && n.Y <= maxy {
						antigrid[n] = true
					}
				}
			}
		}
		for k := range antigrid {
			answer[k] = '#'
		}
	}

	for y := 0; y <= maxy; y++ {
		for x := 0; x <= maxx; x++ {
			r, ok := answer[aoc2024.Coord{X: x, Y: y}]
			if ok {
				fmt.Print(string(r))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	count := 0
	for _, v := range answer {
		if v == '#' {
			count++
		}
	}

	return count
}

func antinodes(a, b aoc2024.Coord, repeat bool) []aoc2024.Coord {
	dir1 := a.Diff(b)
	dir2 := b.Diff(a)
	var result []aoc2024.Coord
	num := 1
	if repeat {
		num = 55
		result = append(result, a)
		result = append(result, b)
	}
	for i := 0; i < num; i++ {
		a = a.Move(dir1)
		b = b.Move(dir2)
		result = append(result, a)
		result = append(result, b)
	}

	return result
}
