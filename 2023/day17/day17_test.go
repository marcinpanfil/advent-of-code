package aoc2023

import (
	"advent-of-code/util/file"
	"fmt"
	"strings"
	"testing"
)

var INPUT = `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533`

func TestSolution(t *testing.T) {
	input := file.ReadInput()
	fmt.Println(FindMinPath(input, 3, 0))  // 886
	fmt.Println(FindMinPath(input, 10, 4)) // 1055
}

func TestFindMinPath(t *testing.T) {
	result1 := FindMinPath(strings.Split(INPUT, "\n"), 3, 0)
	result2 := FindMinPath(strings.Split(INPUT, "\n"), 10, 4)
	if result1 == 102 && result2 == 94 {
		fmt.Println("Correct!")
	} else {
		t.Errorf("Wrong result: %v, %v!\n", result1, result2)
	}
}
