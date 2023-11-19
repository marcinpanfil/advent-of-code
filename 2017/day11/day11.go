package aoc2017

import (
	"advent-of-code/util/math"
	_math "math"
)

var DIRECTIONS = map[string][]int{
	"n":  {0, 2},
	"nw": {-1, 1},
	"sw": {-1, -1},
	"s":  {0, -2},
	"se": {1, -1},
	"ne": {1, 1},
}

func CalculateSteps(input []string) int {
	x := 0
	y := 0
	for _, process := range input {
		dir := DIRECTIONS[process]
		x += dir[0]
		y += dir[1]
	}

	return math.Abs(x) + math.Abs((math.Abs(y)-math.Abs(x)))/2
}

func CalculateFurthest(input []string) int {
	x := 0
	y := 0
	max := _math.MinInt32
	for _, process := range input {
		dir := DIRECTIONS[process]
		x += dir[0]
		y += dir[1]
		dist := math.Abs(x) + math.Abs((math.Abs(y)-math.Abs(x)))/2
		if dist > max {
			max = dist
		}
	}

	return max
}
