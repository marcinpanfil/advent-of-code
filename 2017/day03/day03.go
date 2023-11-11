package aoc2017

import (
	"advent-of-code/util/math"
)

func ConstructSquareSequentially(square int) int {
	prev := 0
	currentSize := 1
	startRange := 0
	endRange := 0
	if square == 1 {
		return 0
	}
	for {
		size := (currentSize * currentSize) - prev
		startRange = endRange + 1
		endRange = startRange + size - 1

		if startRange <= square && square <= endRange {
			// calculate the position of the given square,
			// the smallest distance is in the mid of each side, the biggest distance in the corners
			index := ((math.Abs(startRange - square)) % (currentSize - 1)) + 1
			mid := int(currentSize / 2)
			return mid + math.Abs(index-mid)
		}

		prev += size
		currentSize += 2
	}
}

func ConstructSquareAdjacent(squareValue int) int {
	// x, y
	directions := [][2]int{
		{0, 1},  // up
		{-1, 0}, // left
		{0, -1}, // down
		{1, 0},  // right
	}

	cordinatesToValue := map[[2]int]int{
		{0, 0}: 1,
		{1, 0}: 1,
	}

	currentMax := 1
	lastCordinates := [2]int{1, 0}
	for {
		for _, dir := range directions {
			// if the spiral moves to right, it requires an extra step, which is later on the new max length
			if dir[0] == 1 && dir[1] == 0 {
				currentMax += 1
			}
			for math.Abs(lastCordinates[0]+dir[0]) <= currentMax && math.Abs(lastCordinates[1]+dir[1]) <= currentMax {
				newCordinates := [2]int{lastCordinates[0] + dir[0], lastCordinates[1] + dir[1]}
				sum := sumNeighbours(newCordinates, cordinatesToValue)
				if sum > squareValue {
					return sum
				}
				cordinatesToValue[newCordinates] = sum
				lastCordinates = newCordinates
			}
		}
	}
}

var NEIGHBOURS = [][2]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
	{1, 1},
	{1, -1},
	{-1, 1},
	{-1, -1},
}

func sumNeighbours(newCordinates [2]int, cordinatesToValue map[[2]int]int) int {
	sum := 0
	for _, neighbour := range NEIGHBOURS {
		x := newCordinates[0] + neighbour[0]
		y := newCordinates[1] + neighbour[1]
		val, isPresent := cordinatesToValue[[2]int{x, y}]
		if isPresent {
			sum += val
		}
	}
	return sum
}
