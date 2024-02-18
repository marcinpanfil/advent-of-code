package aoc2017

import (
	"advent-of-code/util/file"
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	input := file.ReadInput()
	particles := ParseInput(input)
	fmt.Println(FindParticleWithMinDist(particles)) // 300
	fmt.Println(Collide(particles))                 // 502
}
