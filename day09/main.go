package main

import (
	"fmt"
	"strings"

	"github.com/srowles/aoc2024"
)

func main() {
	input := aoc2024.InputFromWebsite("9")

	expanded := expand(strings.TrimSpace(input))
	compact(expanded)
	fmt.Println(sum(expanded))

	expanded = expand(strings.TrimSpace(input))
	compact2(expanded)
	fmt.Println(sum(expanded))
}

func expand(input string) []int64 {
	var output []int64
	free := false
	i := int64(0)
	for _, r := range input {
		n := int(aoc2024.Int(string(r)))
		if free {
			for k := 0; k < n; k++ {
				output = append(output, -1)
			}
		} else {
			for k := 0; k < n; k++ {
				output = append(output, i)
			}
			i++
		}
		free = !free
	}
	return output
}

func compact(list []int64) {
	p := 0
	for i := len(list) - 1; i >= 0 && i > p; i-- {
		if list[i] == -1 {
			continue
		}
		for list[p] != -1 {
			p++
			if p == i || p >= len(list) {
				return
			}
		}
		list[p] = list[i]
		list[i] = -1
	}
}

func compact2(list []int64) {
	largestFileID := int64(0)
	for _, id := range list {
		if id > largestFileID {
			largestFileID = id
		}
	}
	for id := largestFileID; id >= 0; id-- {
		start, end := -1, -1
		for i, f := range list {
			if f == id && start == -1 {
				start = i
			}
			if i == len(list)-1 {
				end = len(list)
			}
			if f != id && start != -1 {
				end = i
				break
			}
		}
		length := end - start
		if length == 0 {
			continue
		}
		nstart, nend := -1, -1
		for i, f := range list {
			if i > start {
				break
			}
			if f == -1 && nstart == -1 {
				nstart = i
			}
			if f != -1 && nstart != -1 {
				nend = i - 1
				if nend-nstart+1 >= length {
					break
				} else {
					nstart, nend = -1, -1
				}
			}
		}
		if nstart == -1 || nend == -1 {
			continue
		}
		for i := 0; i < length; i++ {
			list[nstart+i] = list[start+i]
			list[start+i] = -1
		}
	}
}

func sum(list []int64) int64 {
	sum := int64(0)
	for i, v := range list {
		if v == -1 {
			continue
		}
		sum += int64(i) * v
	}
	return sum
}
