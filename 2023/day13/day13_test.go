package aoc2023

import (
	"advent-of-code/util/file"
	"fmt"
	"strings"
	"testing"
)

var INPUT = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

func TestSolution(t *testing.T) {
	input := file.ReadInputAsSingle()
	fmt.Println(FindReflectsWithSmudge(strings.Split(input, "\n\n"), 0)) // 30487
	fmt.Println(FindReflectsWithSmudge(strings.Split(input, "\n\n"), 1)) // 31954
}

func TestFindReflectsWithSmudge(t *testing.T) {
	result := FindReflectsWithSmudge(strings.Split(INPUT, "\n\n"), 1)
	if result == 400 {
		fmt.Println("Correct!")
	} else {
		t.Errorf("Wrong result: %v!", result)
	}
}

func TestFindReflects(t *testing.T) {
	result := FindReflectsWithSmudge(strings.Split(INPUT, "\n\n"), 0)
	if result == 405 {
		fmt.Println("Correct!")
	} else {
		t.Errorf("Wrong result: %v!", result)
	}
}
