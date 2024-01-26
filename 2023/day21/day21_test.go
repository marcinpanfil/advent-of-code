package aoc2023

import (
	"advent-of-code/util/file"
	"fmt"
	"strings"
	"testing"
)

var INPUT = `...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........`

func TestSolution(t *testing.T) {
	input := file.ReadInput()
	fmt.Println(ReachGardenPlots(input, 64)) // 3697
	fmt.Println(InterpolationOfPlots(input)) // 608152828731262
}

func TestReachGardenPlots(t *testing.T) {
	res := ReachGardenPlots(strings.Split(INPUT, "\n"), 6)
	if res == 16 {
		fmt.Println("Correct!")
	} else {
		t.Errorf("Wrong result: %v!\n", res)
	}
}
