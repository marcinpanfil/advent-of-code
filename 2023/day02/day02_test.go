package aoc2023

import (
	"advent-of-code/util/file"
	"fmt"
	"strings"
	"testing"
)

var INPUT = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

func TestSolution(t *testing.T) {
	input := file.ReadInput()
	fmt.Println(SumOfPossibleGames(input))   // 2771
	fmt.Println(FindMinNumberOfCubes(input)) // 70924
}

func TestSumOfPossibleGames(t *testing.T) {
	result := SumOfPossibleGames(strings.Split(INPUT, "\n"))
	if result == 8 {
		fmt.Println("Correct!")
	} else {
		t.Errorf("Wrong value: %v\n", result)
	}
}

func TestFindMinNumberOfCubes(t *testing.T) {
	result := FindMinNumberOfCubes(strings.Split(INPUT, "\n"))
	if result == 2286 {
		fmt.Println("Correct!")
	} else {
		t.Errorf("Wrong value: %v\n", result)
	}
}
