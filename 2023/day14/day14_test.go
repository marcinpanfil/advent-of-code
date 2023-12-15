package aoc2023

import (
	"advent-of-code/util/file"
	"fmt"
	"strings"
	"testing"
)

var INPUT = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

func TestSolution(t *testing.T) {
	input := file.ReadInput()
	fmt.Println(CalculateTotalLoad(input))        // 113486
	fmt.Println(CalculateTotalLoadAfter1T(input)) //104409
}

func TestCalculateTotalLoad(t *testing.T) {
	result := CalculateTotalLoad(strings.Split(INPUT, "\n"))
	if result == 136 {
		fmt.Println("Works!")
	} else {
		t.Errorf("Wrong result: %v!", result)
	}
}

func TestCalculateTotalLoadAfter1T(t *testing.T) {
	result := CalculateTotalLoadAfter1T(strings.Split(INPUT, "\n"))
	if result == 64 {
		fmt.Println("Works!")
	} else {
		t.Errorf("Wrong result: %v!", result)
	}
}
