package aoc2023

import (
	"advent-of-code/util/array"
	"advent-of-code/util/file"
	"fmt"
	"strings"
	"testing"
)

var INPUT_1 = `.....
.S-7.
||-|.
|L-J-
..-..`

var INPUT_2 = `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`

var INPUT_3 = `...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`

var INPUT_4 = `.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...`

var INPUT_5 = `FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`

func TestSolution(t *testing.T) {
	input := file.ReadInput()
	fmt.Println(FindLongestLoop(input)) // 6757
	fmt.Println(CountInsiders(input))   // 523
}

func TestCountInsiders(t *testing.T) {
	data := map[string]int{
		INPUT_1: 1,
		INPUT_2: 1,
		INPUT_3: 4,
		INPUT_4: 8,
		INPUT_5: 10,
	}
	for input, expected := range data {
		result := CountInsiders(strings.Split(input, "\n"))
		if result == expected {
			fmt.Println("Correct!")
		} else {
			t.Errorf("Inccorect result %v, expected %v\n", result, expected)
		}
	}
}

func TestFindLongestLoop(t *testing.T) {
	data := map[string]int{
		INPUT_1: 4,
		INPUT_2: 8,
	}
	for input, expected := range data {
		result := FindLongestLoop(strings.Split(input, "\n"))
		if result == expected {
			fmt.Println("Correct!")
		} else {
			t.Errorf("Inccorect result %v, expected %v\n", result, expected)
		}
	}
}

func TestGetNeighbors(t *testing.T) {
	data := map[[2]int][][2]int{
		{1, 1}: {{1, 2}, {2, 1}},
		{3, 2}: {{3, 3}, {3, 1}},
		{1, 3}: {{1, 2}, {2, 3}},
	}
	for pos, expected := range data {
		result := GetNeighbors(strings.Split(INPUT_1, "\n"), pos)
		if array.PositionsEquals(result, expected) {
			fmt.Println("Correct!")
		} else {
			t.Errorf("Inccorect result %v for %v. Expected %v\n", result, pos, expected)
		}
	}
}
