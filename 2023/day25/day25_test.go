package aoc2023

import (
	"advent-of-code/util/file"
	"fmt"
	"testing"
)

func TestFind3Components(t *testing.T) {
	input := file.ReadInput()
	fmt.Println(Find6WiresToRemove(input)) // 601310
}
