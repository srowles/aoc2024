package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/srowles/aoc2024"
)

func main() {
	input := aoc2024.InputFromWebsite("6")

	fmt.Println(len(part1(input)))
	fmt.Println(part2(input))
}

var (
	dirs = []aoc2024.Coord{
		{X: 0, Y: -1},
		{X: 1, Y: 0},
		{X: 0, Y: 1},
		{X: -1, Y: 0},
	}
)

func part1(input string) map[aoc2024.Coord]bool {
	grid, start, maxx, maxy := makeGrid(input)

	visited := make(map[aoc2024.Coord]bool)
	guard := start
	visited[guard] = true
	d := dirs[0]
	for {
		if grid[guard.Move(d)] == '#' {
			d = nextDir(d)
		}
		guard = guard.Move(d)
		if guard.X < 0 || guard.Y < 0 || guard.X > maxx || guard.Y > maxy {
			break
		}
		visited[guard] = true
	}

	return visited
}

func part2(input string) int {
	grid, start, maxx, maxy := makeGrid(input)

	known := part1(input)
	loopCount := 0

	delete(known, start) // can't be at start
	var wg sync.WaitGroup
	for kn := range known {
		wg.Add(1)
		go func() {
			defer wg.Done()
			d := dirs[0]
			guard := start
			visited := make(map[aoc2024.Coord]int)
			ngrid := make(map[aoc2024.Coord]rune)
			for k, v := range grid {
				ngrid[k] = v
			}
			ngrid[aoc2024.Coord{X: kn.X, Y: kn.Y}] = 'O'
			for s := 0; s < 1000000; s++ {
				if ngrid[guard.Move(d)] == '#' || ngrid[guard.Move(d)] == 'O' {
					d = nextDir(d)
					if visited[guard] >= 8 {
						loopCount++
						break
					}
					continue
				}
				guard = guard.Move(d)
				if guard.X < 0 || guard.Y < 0 || guard.X > maxx || guard.Y > maxy {
					break
				}
				visited[guard] += 1
			}
		}()
	}
	wg.Wait()

	return loopCount
}

func makeGrid(input string) (map[aoc2024.Coord]rune, aoc2024.Coord, int, int) {
	grid := make(map[aoc2024.Coord]rune)
	var start aoc2024.Coord
	var maxx, maxy int
	for y, row := range strings.Split(input, "\n") {
		if y > maxy {
			maxy = y
		}
		for x, ch := range row {
			if x > maxx {
				maxx = x
			}
			if ch == '.' {
				continue
			}
			c := aoc2024.Coord{X: x, Y: y}
			if ch == '^' {
				start = c
			}
			grid[c] = ch
		}
	}
	return grid, start, maxx, maxy
}

func nextDir(dir aoc2024.Coord) aoc2024.Coord {
	for i := 0; i < len(dirs)-1; i++ {
		if dirs[i].Equal(dir) {
			return dirs[i+1]
		}
	}
	return dirs[0]
}
