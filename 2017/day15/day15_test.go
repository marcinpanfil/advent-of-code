package aoc2017

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	fmt.Println(CalculateValueOfGen(873, 583))
	fmt.Println(CalculateValueOfPickyGen(873, 583))
}

func TestCalculateValueOfGen(t *testing.T) {
	result := CalculateValueOfGen(65, 8921)
	if result == 588 {
		fmt.Println("Correct!")
	} else {
		t.Errorf("Incorrect result: %v\n", result)
	}
}

func TestCalculateValueOfPickyGen(t *testing.T) {
	result := CalculateValueOfPickyGen(65, 8921)
	if result == 309 {
		fmt.Println("Correct!")
	} else {
		t.Errorf("Incorrect result: %v\n", result)
	}
}
