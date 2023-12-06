package aoc2023

import (
	"fmt"
	"testing"
)

var INPUT = map[int]int{
	7:  9,
	15: 40,
	30: 200,
}

var SOLUTION_INPUT = map[int]int{
	57: 291,
	72: 1172,
	69: 1176,
	92: 2026,
}

func TestSolution(t *testing.T) {
	fmt.Println(CountNumberOfBestWays(SOLUTION_INPUT))                         // 160816
	fmt.Println(CountNumberOfBestWays(map[int]int{57726992: 291117211762026})) // 46561107
}

func TestCountNumberOfBestWays(t *testing.T) {
	result := CountNumberOfBestWays(INPUT)
	if result == 288 {
		fmt.Println("Works!")
	} else {
		t.Errorf("Wrong value: %v!", result)
	}

}
