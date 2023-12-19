package aoc2023

import (
	"advent-of-code/util/file"
	"fmt"
	"strings"
	"testing"
)

var INPUT = `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

func TestSolution(t *testing.T) {
	input := file.ReadInput()
	fmt.Println(CountAllEnergized(input, -1, 0, [2]int{1, 0})) // 7242
	fmt.Println(CountFromAllTiles(input))                      // 7572
}

func TestCountFromAllTiles(t *testing.T) {
	result := CountFromAllTiles(strings.Split(INPUT, "\n"))
	if result == 51 {
		fmt.Println("Works!")
	} else {
		t.Errorf("Incorrect result: %v!\n", result)
	}
}

func TestCountAllEnergized(t *testing.T) {
	result := CountAllEnergized(strings.Split(INPUT, "\n"), -1, 0, [2]int{1, 0})
	if result == 46 {
		fmt.Println("Works!")
	} else {
		t.Errorf("Incorrect result: %v!\n", result)
	}
}
