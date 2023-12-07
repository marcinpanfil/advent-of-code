package aoc2017

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	fmt.Println(SpinlockAlgorithm(370))     // 1244
	fmt.Println(SpinlockAlgorithm50kk(370)) // 11162912
}

func TestSpinlockAlgorithm(t *testing.T) {
	result := SpinlockAlgorithm(3)
	if result == 638 {
		fmt.Println("Correct!")
	} else {
		t.Errorf("Wrong value %v!", result)
	}
}
