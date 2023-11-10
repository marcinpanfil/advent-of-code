package aoc2017

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	Part1()
	Part2()
}

func TestCalculateSpreadsheetChecksum(t *testing.T) {
	data := [][]int{
		{5, 1, 9, 5},
		{7, 5, 3},
		{2, 4, 6, 8},
	}
	if CalculateSpreadsheetChecksum(data) == 18 {
		fmt.Println("Works!")
	} else {
		t.Error("Does not work!")
	}
}

func TestCalculateSpreadsheetChecksumUsingDiv(t *testing.T) {
	data := [][]int{
		{5, 9, 2, 8},
		{9, 4, 7, 3},
		{3, 8, 6, 5},
	}

	if CalculateSpreadsheetChecksumUsingDiv(data) == 9 {
		fmt.Println("Works!")
	} else {
		t.Error("Does not work!")
	}
}
