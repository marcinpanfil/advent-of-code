package aoc2017

import (
	"advent-of-code/util/file"
	"fmt"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	fmt.Println(FindTheLargestRegister(file.ReadInput())) //8022 9819
}

var INPUT string = `b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10`

func TestFindTheLargestRegister(t *testing.T) {
	max, historialMax := FindTheLargestRegister(strings.Split(INPUT, "\n"))
	if max == 1 && historialMax == 10 {
		fmt.Println("Correct!")
	} else {
		t.Errorf("The %v and %v values are incorrect!", max, historialMax)
	}
}
