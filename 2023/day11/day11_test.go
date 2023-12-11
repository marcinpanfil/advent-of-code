package aoc2023

import (
	"advent-of-code/util/file"
	"fmt"
	"strings"
	"testing"
)

var INPUT = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

func TestSolution(t *testing.T) {
	input := file.ReadInput()
	fmt.Println(LenghtsBetweenGalaxiesWithXFactor(input, 2))         // 9233514
	fmt.Println(LenghtsBetweenGalaxiesWithXFactor(input, 1_000_000)) // 363293506944
}

func TestLenghtsBetweenGalaxiesWithXFactor(t *testing.T) {
	data := map[int]int{
		2:   374,
		10:  1030,
		100: 8410,
	}
	for factor, expected := range data {
		result := LenghtsBetweenGalaxiesWithXFactor(strings.Split(INPUT, "\n"), factor)
		if result == expected {
			fmt.Println("Correct!")
		} else {
			t.Errorf("Incorrect value %v, expected %v\n", result, expected)
		}
	}
}
