package aoc2017

import (
	"advent-of-code/util/file"
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	input := file.ReadInput()
	fmt.Println(FindTotalScore(input[0]))          // 11347
	fmt.Println(CountCharsWithinGarbage(input[0])) // 5404
}

func TestCountCharsWithinGarbage(t *testing.T) {
	data := map[string]int{
		"<>":                  0,
		"<random characters>": 17,
		"<<<<>":               3,
		"<{!>}>":              2,
		"<!!>":                0,
		"<!!!>>":              0,
		"<{o\"i!a,<{i<a>":     10,
	}

	for k, v := range data {
		result := CountCharsWithinGarbage(k)
		if v == result {
			fmt.Printf("For value %v there is correct result %v\n", k, v)
		} else {
			t.Errorf("For value %v there is incorrect result %v\n", k, result)
		}
	}
}

func TestFindTotalScore(t *testing.T) {
	data := map[string]int{
		"{}":                            1,
		"{{{}}}":                        6,
		"{{},{}}":                       5,
		"{{{},{},{{}}}}":                16,
		"{<a>,<a>,<a>,<a>}":             1,
		"{{<ab>},{<ab>},{<ab>},{<ab>}}": 9,
		"{{<!!>},{<!!>},{<!!>},{<!!>}}": 9,
		"{{<a!>},{<a!>},{<a!>},{<ab>}}": 3,
	}

	for k, v := range data {
		result := FindTotalScore(k)
		if v == result {
			fmt.Printf("For value %v there is correct result %v\n", k, v)
		} else {
			t.Errorf("For value %v there is incorrect result %v\n", k, result)
		}
	}

}
