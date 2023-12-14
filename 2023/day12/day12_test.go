package aoc2023

import (
	"advent-of-code/util/file"
	"fmt"
	"strings"
	"testing"
)

var INPUT = `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`

func TestSolution(t *testing.T) {
	input := file.ReadInput()
	fmt.Println(CountPossibleArrangement(input, 1)) // 7407
	fmt.Println(CountPossibleArrangement(input, 5)) // 30568243604962
}

func TestPermutation(t *testing.T) {
	data := map[string]int{
		"?.# 1,1":                   1,
		"???.### 1,1,3":             1,
		".??..??...?##. 1,1,3":      4,
		"?#?#?#?#?#?#?#? 1,3,1,6":   1,
		"????.#...#... 4,1,1":       1,
		"????.######..#####. 1,6,5": 4,
		"?###???????? 3,2,1":        10,
	}
	for input, expected := range data {
		result := CountPossibleArrangement([]string{input}, 1)
		if result == expected {
			fmt.Printf("For %s it works!\n", input)
		} else {
			t.Errorf("For %s wrong result: %v, expected %v\n", input, result, expected)
		}
	}
}

func TestFindPossibleArrangementsSingleLineWithFactor(t *testing.T) {
	data := map[string]int{
		"???.### 1,1,3":             1,
		".??..??...?##. 1,1,3":      16384,
		"?#?#?#?#?#?#?#? 1,3,1,6":   1,
		"????.#...#... 4,1,1":       16,
		"????.######..#####. 1,6,5": 2500,
		"?###???????? 3,2,1":        506250,
	}
	for input, expected := range data {
		result := CountPossibleArrangement([]string{input}, 5)
		if result == expected {
			fmt.Printf("For %s it works!\n", input)
		} else {
			t.Errorf("For %s wrong result: %v, expected %v\n", input, result, expected)
		}
	}
	res := CountPossibleArrangement(strings.Split(INPUT, "\n"), 5)
	if res == 525152 {
		fmt.Printf("For the general it works!\n")
	} else {
		t.Errorf("For the general wrong result: %v, expected 21\n", res)
	}
}
