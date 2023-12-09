package aoc2023

import (
	"advent-of-code/util/file"
	"fmt"
	"strings"
	"testing"
)

var INPUT = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

func TestSolution(t *testing.T) {
	input := file.ReadInput()
	fmt.Println(ExtrapolateValues(input)) // 1772145754, 867
}

func TestExtrapolateValues(t *testing.T) {
	result1, result2 := ExtrapolateValues(strings.Split(INPUT, "\n"))
	if result1 == 114 && result2 == 2 {
		fmt.Println("Correct!")
	} else {
		t.Errorf("Incorrect result %v %v!", result1, result2)
	}
}
