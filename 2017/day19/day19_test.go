package aoc2017

import (
	"advent-of-code/util/file"
	"fmt"
	"strings"
	"testing"
)

var INPUT = `     |          
     |  +--+    
     A  |  C    
 F---|----E|--+ 
     |  |  |  D 
     +B-+  +--+ 
`

func TestSolution(t *testing.T) {
	input := file.ReadInput()
	fmt.Println(FindPath(input)) // KGPTMEJVS, 16328
}

func TestFindPath(t *testing.T) {
	path, steps := FindPath(strings.Split(INPUT, "\n"))
	if path == "ABCDEF" && steps == 38 {
		fmt.Println("Correct!")
	} else {
		t.Errorf("Wrong result: %s %v\n", path, steps)
	}
}
