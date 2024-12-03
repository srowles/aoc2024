package main

import (
	"fmt"

	"github.com/srowles/aoc2024"
)

func main() {
	input := aoc2024.InputFromWebsite("3")

	fmt.Println(mul(input, false))
	fmt.Println(mul(input, true))
}

type instr struct {
	a int
	b int
}

func mul(line string, does bool) int {
	result := 0
	for _, v := range muls(line, does) {
		result += v.a * v.b
	}
	return result
}

func muls(line string, does bool) []instr {
	var result []instr
	do := true
	for i := 0; i < len(line); i++ {
		check := line[i:]
		if does && len(check) > 4 && check[:4] == "do()" {
			do = true
			continue
		}
		if does && len(check) > 7 && check[:7] == "don't()" {
			do = false
			continue
		}
		var a, b int
		n, err := fmt.Sscanf(check, "mul(%d,%d)", &a, &b)
		if n != 2 || err != nil {
			continue
		}
		if do {
			result = append(result, instr{a: a, b: b})
		}
	}

	return result
}
