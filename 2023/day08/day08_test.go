package aoc2023

import (
	"advent-of-code/util/file"
	"fmt"
	"strings"
	"testing"
)

var INPUT_1 = `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`

var INPUT_2 = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

var INPUT_3 = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

func TestSolution(t *testing.T) {
	input := file.ReadInput()
	fmt.Println(CalculateSteps(input))               // 17287
	fmt.Println(CalculateStepsSimultaneously(input)) // 18625484023687
}

func TestCalculateSteps(t *testing.T) {
	data := map[string]int{
		INPUT_1: 2,
		INPUT_2: 6,
	}
	for k, v := range data {
		result := CalculateSteps(strings.Split(k, "\n"))
		if result == v {
			fmt.Println("Works!")
		} else {
			t.Errorf("For %s there's wrong number of steps: %v\n", k, result)
		}
	}
}

func TestCalculateStepsSimultaneously(t *testing.T) {
	result := CalculateStepsSimultaneously(strings.Split(INPUT_3, "\n"))
	if result == 6 {
		fmt.Println("Works!")
	} else {
		t.Errorf("There's wrong number of steps: %v\n", result)
	}
}
