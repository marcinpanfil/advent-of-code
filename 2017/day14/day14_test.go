package aoc2017

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	fmt.Printf("%v\n", CalculateUsed("ljoxqyyw"))   // 8316
	fmt.Printf("%v\n", CalculateGroups("ljoxqyyw")) // 1074
}

func TestCalculateGroups(t *testing.T) {
	result := CalculateGroups("flqrgnkx")
	if result == 1242 {
		fmt.Println("Correct!")
	} else {
		t.Errorf("Incorrect result %v\n", result)
	}
}

func TestStringToBin(t *testing.T) {
	result := CalculateUsed("flqrgnkx")
	if result == 8108 {
		fmt.Println("Correct!")
	} else {
		t.Errorf("Incorrect result %v\n", result)
	}
}
