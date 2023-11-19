package aoc2017

import (
	"advent-of-code/util/file"
	"fmt"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	input := strings.Split(file.ReadInput()[0], ",")
	fmt.Println(CalculateSteps(input))    // 682
	fmt.Println(CalculateFurthest(input)) // 1406
}

func TestCalculateSteps(t *testing.T) {
	data := map[string]int{
		"ne,ne,ne":       3,
		"ne,ne,sw,sw":    0,
		"ne,ne,s,s":      2,
		"se,sw,se,sw,sw": 3,
	}
	for k, v := range data {
		result := CalculateSteps(strings.Split(k, ","))
		if result == v {
			fmt.Printf("For %s the value is correct!\n", k)
		} else {
			t.Errorf("For %s the value %v is incorrect!\n", k, result)
		}
	}
}
