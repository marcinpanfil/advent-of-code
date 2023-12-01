package aoc2023

import (
	"advent-of-code/util/file"
	"fmt"
	"strings"
	"testing"
)

var INPUT_1 = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

var INPUT_2 = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

func TestSolution(t *testing.T) {
	data := file.ReadInput()
	fmt.Println(SumCalibraionValue(data))            // 56506
	fmt.Println(SumCalibraionValueWithLetters(data)) // 56017
}

func TestSumCalibraionValue(t *testing.T) {
	result := SumCalibraionValue(strings.Split(INPUT_1, "\n"))
	if result == 142 {
		fmt.Println("Correct!")
	} else {
		t.Errorf("Wrong value: %v\n", result)
	}
}

func TestSumCalibraionValueWithLetters(t *testing.T) {
	result := SumCalibraionValueWithLetters(strings.Split(INPUT_2, "\n"))
	if result == 281 {
		fmt.Println("Correct!")
	} else {
		t.Errorf("Wrong value: %v\n", result)
	}
}
