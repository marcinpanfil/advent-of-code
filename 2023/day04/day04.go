package aoc2023

import (
	"advent-of-code/util/array"
	"math"
	"strings"
)

func CalculateValueOfScratchcards(input []string) int {
	result := 0
	for _, line := range input {
		count := 0
		withName := strings.Split(line, ": ")
		allNumbersAsStr := strings.Split(withName[1], " | ")
		myNumbersAsStr := strings.Fields(strings.Trim(allNumbersAsStr[0], " "))
		winningNumbersAsStr := strings.Fields(strings.Trim(allNumbersAsStr[1], " "))
		for _, number := range myNumbersAsStr {
			if array.IndexOf[string](winningNumbersAsStr, number) >= 0 {
				count++
			}
		}
		if count > 0 {
			result += 1 * int(math.Pow(2, float64(count-1)))
		}
	}
	return result
}

func CalculateValueOfScratchcardsWithCopies(input []string) int {
	result := 0
	additionalCopies := map[int]int{}
	for idx, line := range input {
		count := 0
		withName := strings.Split(line, ": ")
		allNumbersAsStr := strings.Split(withName[1], " | ")
		myNumbersAsStr := strings.Fields(strings.Trim(allNumbersAsStr[0], " "))
		winningNumbersAsStr := strings.Fields(strings.Trim(allNumbersAsStr[1], " "))
		for _, number := range myNumbersAsStr {
			if array.IndexOf[string](winningNumbersAsStr, number) >= 0 {
				count++
			}
		}
		copies := additionalCopies[idx]
		for i := idx + 1; i <= idx+count; i++ {
			additionalCopies[i] = additionalCopies[i] + copies + 1
		}
		additionalCopies[idx] = additionalCopies[idx] + 1
	}
	for _, v := range additionalCopies {
		result += v
	}
	return result
}
