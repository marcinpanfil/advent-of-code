package aoc2017

import (
	"advent-of-code/util/file"
	"fmt"
	"strings"
	"testing"
)

var INPUT = `0 <-> 2
1 <-> 1
2 <-> 0, 3, 4
3 <-> 2, 4
4 <-> 2, 3, 6
5 <-> 6
6 <-> 4, 5`

func TestSolution(t *testing.T) {
	input := file.ReadInput()
	fmt.Println(CountProgramsConnectedTo0(input)) // 134
	fmt.Println(CountGroups(input))               // 193
}

func TestCountProgramsConnectedTo0(t *testing.T) {
	if CountProgramsConnectedTo0(strings.Split(INPUT, "\n")) == 6 {
		fmt.Println("Correct!")
	} else {
		t.Error("Incorrect!\n")
	}
}

func TestCountGroups(t *testing.T) {
	if CountGroups(strings.Split(INPUT, "\n")) == 2 {
		fmt.Println("Correct!")
	} else {
		t.Error("Incorrect!\n")
	}
}
