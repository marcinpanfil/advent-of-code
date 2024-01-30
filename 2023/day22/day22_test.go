package aoc2023

import (
	"advent-of-code/util/file"
	"fmt"
	"strings"
	"testing"
)

var INPUT = `1,0,1~1,2,1
0,0,2~2,0,2
0,2,3~2,2,3
0,0,4~0,2,4
2,0,5~2,2,5
0,1,6~2,1,6
1,1,8~1,1,9`

var INPUT_1 = `0,0,1~0,1,1
1,1,1~1,1,1
0,0,2~0,0,2
0,1,2~1,1,2`

var INPUT_2 = `0,0,1~1,0,1
0,1,1~0,1,2
0,0,5~0,0,5
0,0,4~0,1,4`

var INPUT_3 = `1,0,1~1,2,1
1,1,2~1,3,2`

func TestSolution(t *testing.T) {
	input := file.ReadInput()
	fmt.Println(SolutionPart1(input)) // 437
	fmt.Println(SolutionPart2(input)) // 42561
}

func TestSolutionPart1(t *testing.T) {
	data := map[string]int{
		INPUT:   5,
		INPUT_1: 3,
		INPUT_2: 2,
		INPUT_3: 1,
	}
	for testCase, expectedResult := range data {
		count := SolutionPart1(strings.Split(testCase, "\n"))
		if count == expectedResult {
			fmt.Println("Correct!")
		} else {
			t.Errorf("Wrong result: %v, expected: %v\n", count, expectedResult)
		}
	}
}

func TestFindSumOfFallenBricks(t *testing.T) {
	count := SolutionPart2(strings.Split(INPUT, "\n"))
	if count == 7 {
		fmt.Println("Works!")
	} else {
		t.Errorf("Wrong result: %v! Expected %v.\n", count, 7)
	}
}
