package aoc2017

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	SolutionPart1() // 387096
	SolutionPart2() // 28040648
}

func TestCalculateJumpsUntilExit(t *testing.T) {
	result := CalculateJumpsUntilExit([]int{0, 3, 0, 1, -3})
	if result == 5 {
		fmt.Println("Works!")
	} else {
		t.Errorf("Does not work! Result %v", result)
	}
}

func TestCalculateJumpsUntilExitWithDecrease(t *testing.T) {
	result := CalculateJumpsUntilExitWithDecrease([]int{0, 3, 0, 1, -3})
	if result == 10 {
		fmt.Println("Works!")
	} else {
		t.Errorf("Does not work! Result %v", result)
	}
}
