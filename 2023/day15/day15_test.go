package aoc2023

import (
	"advent-of-code/util/file"
	"fmt"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	input := file.ReadInputAsSingle()
	fmt.Println(CalculateHashValue(strings.Split(input, ","))) // 509784
	fmt.Println(FillLensSlots(strings.Split(input, ",")))      // 230197
}

func TestCalculateHashValue(t *testing.T) {
	result := CalculateHashValue(strings.Split("rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7", ","))
	if result == 1320 {
		fmt.Println("Correct!")
	} else {
		t.Errorf("Wrong value: %v", result)
	}
}

func TestFillLensSlots(t *testing.T) {
	result := FillLensSlots(strings.Split("rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7", ","))
	if result == 145 {
		fmt.Println("Correct!")
	} else {
		t.Errorf("Wrong value: %v", result)
	}
}
