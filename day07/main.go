package main

import (
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/srowles/aoc2024"
)

func main() {
	input := aoc2024.InputFromWebsite("7")

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int64 {
	total := int64(0)

	for _, line := range aoc2024.Lines(input) {
		answer, valueString, _ := strings.Cut(line, ": ")
		vals := aoc2024.Slice(valueString, " ", aoc2024.Int)
		if valid(aoc2024.Int(answer), vals, 2) {
			total += aoc2024.Int(answer)
		}
	}

	return total
}

func part2(input string) int64 {
	total := int64(0)

	for _, line := range aoc2024.Lines(input) {
		answer, valueString, _ := strings.Cut(line, ": ")
		vals := aoc2024.Slice(valueString, " ", aoc2024.Int)
		if valid(aoc2024.Int(answer), vals, 3) {
			total += aoc2024.Int(answer)
		}
	}

	return total
}

func valid(answer int64, vals []int64, numberOfOps int) bool {
	possibleOps := ops(numberOfOps, len(vals)-1)
	for _, ops := range possibleOps {
		sum := vals[0]
		for i := 0; i < len(vals)-1; i++ {
			sum = operations[ops[i]](sum, vals[i+1])
		}
		if sum == answer {
			return true
		}
	}
	return false
}

var operations = []func(a, b int64) int64{
	mul, add, concat,
}

var mul = func(a, b int64) int64 {
	return a * b
}

var add = func(a, b int64) int64 {
	return a + b
}

var concat = func(a, b int64) int64 {
	return aoc2024.Int(fmt.Sprintf("%d%d", a, b))
}

func ops(k, n int) [][]int64 {
	var ret [][]int64
	v := 0
	for i := 0; i < int(math.Pow(float64(k), float64(n))); i++ {
		strval := big.NewInt(int64(v)).Text(k)
		l := len(strval)
		for j := 0; j < n-l; j++ {
			strval = "0" + strval
		}
		ret = append(ret, aoc2024.Slice(strval, "", aoc2024.Int))
		v += 1
	}
	return ret
}
