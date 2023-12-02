package aoc2017

import (
	"advent-of-code/util/file"
	"fmt"
	"strings"
	"testing"
)

var PROGRAMS = []string{"a", "b", "c", "d", "e"}
var MOVES = []string{"s1", "x3/4", "pe/b"}

func TestSolution(t *testing.T) {
	moves := strings.Split(file.ReadInput()[0], ",")
	fmt.Println(Dance([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"}, moves, 1)) //glnacbhedpfjkiom
	// there is a repeating pattern of dance. In this case after 60 dances, it starts from dance nr 0, 1000 000 000/60= 40
	fmt.Println(Dance([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"}, moves, 40)) //fmpanloehgkdcbji
}

func TestDance(t *testing.T) {
	result := Dance(PROGRAMS, MOVES, 1)
	if result == "baedc" {
		fmt.Println("Correct!")
	} else {
		t.Errorf("Incorrect order: %s", result)
	}
}
