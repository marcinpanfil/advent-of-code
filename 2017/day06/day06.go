package aoc2017

import (
	"advent-of-code/util/array"
	"fmt"
	"strings"
)

func Redistribute(input []int, part1 bool) int {
	steps := 0
	history := map[string]int{}
	for {
		index, maxValue := array.MaxIndexAndValue(input)
		input[index] = 0
		startPos := (index + 1) % len(input)
		for i := startPos; i < startPos+maxValue; i++ {
			input[i%len(input)] = input[i%len(input)] + 1
		}
		steps++
		inputAsStr := ArrayToString(input)
		v, k := history[inputAsStr]
		if k {
			if part1 {
				return steps
			} else {
				return steps - v
			}
		} else {
			history[inputAsStr] = steps
		}
	}
}

func ArrayToString(input []int) string {
	return strings.Trim(strings.Join(strings.Split(fmt.Sprint(input), " "), "|"), "[]")
}
