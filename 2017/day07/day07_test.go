package aoc2017

import (
	"advent-of-code/util/file"
	"fmt"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	fmt.Println(FindBottomProgram(file.ReadInput()))                     // rqwgj
	fmt.Println(FindCorrectWeightOfUnballancedProgram(file.ReadInput())) //333
}

var INPUT string = `pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`

func TestFindBottomProgram(t *testing.T) {
	bottomProgram := FindBottomProgram(strings.Split(INPUT, "\n"))
	if bottomProgram == "tknk" {
		fmt.Println("Correct value!")
	} else {
		t.Errorf("Received wrong result: %s\n", bottomProgram)
	}
}

func TestFindCorrectWeightOfUnballancedProgram(t *testing.T) {
	bottomProgram := FindCorrectWeightOfUnballancedProgram(strings.Split(INPUT, "\n"))
	if bottomProgram == 60 {
		fmt.Println("Correct value!")
	} else {
		t.Errorf("Received wrong result: %v\n", bottomProgram)
	}
}
