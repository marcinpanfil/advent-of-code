package aoc2023

import (
	"advent-of-code/util/file"
	"fmt"
	"strings"
	"testing"
)

var INPUT = `#.#####################
#.......#########...###
#######.#########.#.###
###.....#.>.>.###.#.###
###v#####.#v#.###.#.###
###.>...#.#.#.....#...#
###v###.#.#.#########.#
###...#.#.#.......#...#
#####.#.#.#######.#.###
#.....#.#.#.......#...#
#.#####.#.#.#########v#
#.#...#...#...###...>.#
#.#.#v#######v###.###v#
#...#.>.#...>.>.#.###.#
#####v#.#.###v#.#.###.#
#.....#...#...#.#.#...#
#.#########.###.#.#.###
#...###...#...#...#.###
###.###.#.###v#####v###
#...#...#.#.>.>.#.>.###
#.###.###.#.###.#.#v###
#.....###...###...#...#
#####################.#`

func TestSolution(t *testing.T) {
	input := file.ReadInput()
	fmt.Println(FindTheLongestPath(input))                  // 2414
	fmt.Println(FindTheLongestDistanceWithoutSlopes(input)) // 6598
}

func TestFindTheLongestPath(t *testing.T) {
	result := FindTheLongestPath(strings.Split(INPUT, "\n"))
	if result == 94 {
		fmt.Println("Correct!")
	} else {
		t.Errorf("Wrong result: %v\n", result)
	}
}

func TestCreateMatixOfDistances(t *testing.T) {
	result := FindTheLongestDistanceWithoutSlopes(strings.Split(INPUT, "\n"))
	if result == 154 {
		fmt.Println("Correct!")
	} else {
		t.Errorf("Wrong result: %v\n", result)
	}
}
