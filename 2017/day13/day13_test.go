package aoc2017

import (
	"advent-of-code/util/file"
	"fmt"
	"strings"
	"testing"
)

var INPUT = `0: 3
1: 2
4: 4
6: 4`

func TestSolution(t *testing.T) {
	data := ParseInput(file.ReadInput())
	fmt.Println(SecurityScanner(data)) // 1900
	fmt.Println(FindMaxDelay(data))    // 3966414
}

func TestFindMaxDelay(t *testing.T) {
	data := ParseInput(strings.Split(INPUT, "\n"))
	result := FindMaxDelay(data)
	if result == 10 {
		fmt.Println("Works!")
	} else {
		t.Errorf("Incorrect value %v!\n", result)
	}
}

func TestSecurityScanner(t *testing.T) {
	data := ParseInput(strings.Split(INPUT, "\n"))
	result := SecurityScanner(data)
	if result == 24 {
		fmt.Println("Works!")
	} else {
		t.Errorf("Incorrect value %v!\n", result)
	}
}
