package aoc2023

import (
	"fmt"
	"strconv"
)

func SumCalibraionValue(input []string) int {
	result := 0
	for _, line := range input {
		firstValue := 0
		lastValue := 0
		for _, char := range line {
			if v, err := strconv.Atoi(string(char)); err == nil {
				if firstValue == 0 {
					firstValue = v
				}
				lastValue = v
			}
		}
		calibrationValue, err := strconv.Atoi(fmt.Sprint(firstValue) + fmt.Sprint(lastValue))
		if err == nil {
			result += calibrationValue
		} else {
			fmt.Printf("Wrong value for %v, %v\n", firstValue, lastValue)
			panic("Wrong value!")
		}
	}
	return result
}

func SumCalibraionValueWithLetters(input []string) int {
	mapping := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	result := 0
	for _, line := range input {
		firstValue := 0
		lastValue := 0
		for index, char := range line {
			if v, err := strconv.Atoi(string(char)); err == nil {
				if firstValue == 0 {
					firstValue = v
				}
				lastValue = v
			} else {
				for _, i := range []int{2, 3, 4} {
					if index >= i {
						value, ok := mapping[line[index-i:index+1]]
						if ok {
							if firstValue == 0 {
								firstValue = value
							}
							lastValue = value
						}
					}
				}
			}
		}
		calibrationValue, err := strconv.Atoi(fmt.Sprint(firstValue) + fmt.Sprint(lastValue))
		if err == nil {
			result += calibrationValue
		} else {
			fmt.Printf("Wrong value for %v, %v\n", firstValue, lastValue)
			panic("Wrong value!")
		}
	}
	return result
}
