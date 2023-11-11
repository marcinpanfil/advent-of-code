package aoc2017

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	input := []int{10, 3, 15, 10, 5, 15, 5, 15, 9, 2, 5, 8, 5, 2, 3, 6}
	fmt.Println(Redistribute(input, true))  // 14029
	fmt.Println(Redistribute(input, false)) // 2765
}

func TestRedistribute(t *testing.T) {
	result := Redistribute([]int{0, 2, 7, 0}, true)
	if result == 5 {
		fmt.Println("Part1 works!")
	} else {
		t.Errorf("Does not work for part1! Result: %v", result)
	}
	result = Redistribute([]int{0, 2, 7, 0}, false)
	if result == 4 {
		fmt.Println("Part2 works!")
	} else {
		t.Errorf("Does not work for part2! Result: %v", result)
	}
}
