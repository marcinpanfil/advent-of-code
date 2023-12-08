package aoc2017

import (
	"advent-of-code/util/file"
	"fmt"
	"strings"
	"testing"
)

var INPUT = `set a 1
add a 2
mul a a
mod a 5
snd a
set a 0
rcv a
jgz a -1
set a 1
jgz a -2`

func TestSolution(t *testing.T) {
	input := file.ReadInput()
	fmt.Println(RecoverLastSound(input)) // 7071
	fmt.Println(Start(input))            // 8001
}

func TestRecoverLastSound(t *testing.T) {
	result := RecoverLastSound(strings.Split(INPUT, "\n"))
	if result == 4 {
		fmt.Println("Correct!")
	} else {
		t.Errorf("Wrong result %v!", result)
	}
}

var INPUT1 = `snd 1
snd 2
snd p
rcv a
rcv b
rcv c
rcv d`

func TestStart(t *testing.T) {
	result := Start(strings.Split(INPUT1, "\n"))
	if result == 3 {
		fmt.Println("Works!")
	} else {
		t.Errorf("Wrong value %v\n", result)
	}
}
