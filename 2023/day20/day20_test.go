package aoc2023

import (
	"advent-of-code/util/file"
	"fmt"
	"strings"
	"testing"
)

var INPUT_1 = `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`

var INPUT_2 = `broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output`

func TestSolution(t *testing.T) {
	input := file.ReadInput()
	fmt.Println(PushButton(input))        // 747304011
	fmt.Println(CountButtonPushes(input)) // 220366255099387
}

func TestPushButton(t *testing.T) {
	result1 := PushButton(strings.Split(INPUT_1, "\n"))
	result2 := PushButton(strings.Split(INPUT_2, "\n"))
	if result1 == 32000000 {
		fmt.Println("Correct for result 1!")
	} else {
		t.Errorf("Incorrect result: %v\n", result1)
	}
	if result2 == 11687500 {
		fmt.Println("Correct for result 2!")
	} else {
		t.Errorf("Incorrect result: %v\n", result2)
	}
}
