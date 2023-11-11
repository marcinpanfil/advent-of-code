package aoc2017

import (
	"advent-of-code/util/file"
	"fmt"
)

func SolutionPart1() {
	input := file.ReadInputAsIntArray()
	fmt.Printf("Result part1: %v\n", CalculateJumpsUntilExit(input))
}

func SolutionPart2() {
	input := file.ReadInputAsIntArray()
	fmt.Printf("Result part2: %v\n", CalculateJumpsUntilExitWithDecrease(input))
}

func CalculateJumpsUntilExit(input []int) int {
	currIndex := 0
	jumpsCount := 0
	for {
		currValue := input[currIndex]
		input[currIndex] = input[currIndex] + 1
		currIndex += currValue
		jumpsCount++
		if currIndex >= len(input) || currIndex < 0 {
			return jumpsCount
		}
	}
}

func CalculateJumpsUntilExitWithDecrease(input []int) int {
	currIndex := 0
	jumpsCount := 0
	for {
		currValue := input[currIndex]
		if currValue >= 3 {
			input[currIndex] = input[currIndex] - 1
		} else {
			input[currIndex] = input[currIndex] + 1
		}
		currIndex += currValue
		jumpsCount++
		if currIndex >= len(input) || currIndex < 0 {
			return jumpsCount
		}
	}
}
