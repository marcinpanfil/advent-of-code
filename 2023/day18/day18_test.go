package aoc2023

import (
	"advent-of-code/util/file"
	"fmt"
	"strings"
	"testing"
)

var INPUT = `R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)`

func TestSolution(t *testing.T) {
	input := file.ReadInput()
	fmt.Println(DigInterior(input))        // 61661
	fmt.Println(DigInteriorWithHex(input)) // 111131796939729
}

func TestDigInterior(t *testing.T) {
	result := DigInterior(strings.Split(INPUT, "\n"))
	if result == 62 {
		fmt.Println("Correct!")
	} else {
		t.Errorf("Wrong result %v\n", result)
	}
}

func TestDigInteriorWithHex(t *testing.T) {
	result := DigInteriorWithHex(strings.Split(INPUT, "\n"))
	if result == 952408144115 {
		fmt.Println("Correct!")
	} else {
		t.Errorf("Wrong result %v\n", result)
	}
}
