package aoc2017

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	Part1()
	Part2()
}

func TestSumUpIfXAheadIsEqual(t *testing.T) {
	data := []struct {
		input          string
		expectedResult int
		steps          int
	}{
		{"1122", 3, 1},
		{"1111", 4, 1},
		{"1234", 0, 1},
		{"91212129", 9, 1},
		{"1212", 6, 0},
		{"1221", 0, 0},
		{"123425", 4, 0},
		{"123123", 12, 0},
		{"12131415", 4, 0},
	}

	for _, testCase := range data {
		result := SumUpIfXAheadIsEqual(testCase.input, testCase.steps)
		if testCase.expectedResult == result {
			fmt.Printf("%s works!\n", testCase.input)
		} else {
			t.Errorf("%s does not work. Got %v, expected %v", testCase.input, result, testCase.expectedResult)
		}
	}
}
