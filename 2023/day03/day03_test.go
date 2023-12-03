package aoc2023

import (
	"advent-of-code/util/file"
	"fmt"
	"strings"
	"testing"
)

var INPUT = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func TestSolution(t *testing.T) {
	input := file.ReadInput()
	fmt.Println(CalcEngine(input)) // 527144
	fmt.Println(CalcGear(input))   // 81463996
}

func TestCalcEngine(t *testing.T) {
	result := CalcEngine(strings.Split(INPUT, "\n"))
	if result == 4361 {
		fmt.Println("Correct!")
	} else {
		t.Errorf("Wrong value: %v", result)
	}
}

func TestCalcGear(t *testing.T) {
	result := CalcGear(strings.Split(INPUT, "\n"))
	if result == 467835 {
		fmt.Println("Correct!")
	} else {
		t.Errorf("Wrong value: %v", result)
	}
}
