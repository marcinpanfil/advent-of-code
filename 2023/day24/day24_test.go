package aoc2023

import (
	"advent-of-code/util/file"
	"fmt"
	"strings"
	"testing"
)

var INPUT = `19, 13, 30 @ -2, 1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @ 1, -5, -3`

func TestSolution(t *testing.T) {
	input := file.ReadInput()
	fmt.Println(CalculateCollidePositions(input)) // 13149
	fmt.Println(FindSingleThrow(input))           //1033770143421619
}

func TestCalculateCollidePositions(t *testing.T) {
	MIN = 7
	MAX = 27
	res := CalculateCollidePositions(strings.Split(INPUT, "\n"))
	if res == 2 {
		fmt.Println("Correct!")
	} else {
		t.Errorf("Wrong result %v!", res)
	}
}

func TestFindSingleThrow(t *testing.T) {
	res := FindSingleThrow(strings.Split(INPUT, "\n"))
	if res == 47 {
		fmt.Println("Correct!")
	} else {
		t.Errorf("Wrong result %v!\n", res)
	}
}
