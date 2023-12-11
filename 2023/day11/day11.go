package aoc2023

import (
	"advent-of-code/util/array"
	"advent-of-code/util/math"
)

func LenghtsBetweenGalaxiesWithXFactor(input []string, factor int) int {
	sum := 0
	positions, emptyRows, emptyCols := GetGalaxiesData(input)
	for i := 0; i < len(positions)-1; i++ {
		for j := i + 1; j < len(positions); j++ {
			posA := positions[i]
			posB := positions[j]
			xExpansion := CalculateExpansion(emptyCols, posA[0], posB[0]) * (factor - 1)
			yExpansion := CalculateExpansion(emptyRows, posA[1], posB[1]) * (factor - 1)
			sum += math.Abs(posA[0]-posB[0]) + xExpansion + math.Abs(posA[1]-posB[1]) + yExpansion
		}

	}
	return sum
}

func CalculateExpansion(empty []int, pos1, pos2 int) int {
	start := min(pos1, pos2)
	end := max(pos1, pos2)
	result := 0
	for _, e := range empty {
		if e > start && e < end {
			result++
		}
	}
	return result
}

func GetGalaxiesData(input []string) ([][2]int, []int, []int) {
	pos := [][2]int{}
	emptyRows := []int{}
	hasGalaxy := map[int]bool{}
	for id, row := range input {
		colIds := array.IndexesOf[rune]([]rune(row), '#')
		if len(colIds) > 0 {
			for _, colId := range colIds {
				pos = append(pos, [2]int{colId, id})
				hasGalaxy[colId] = true
			}
		} else {
			emptyRows = append(emptyRows, id)
		}
	}
	emptyCols := []int{}
	for i := 0; i < len(input[0]); i++ {
		if !hasGalaxy[i] {
			emptyCols = append(emptyCols, i)
		}
	}
	return pos, emptyRows, emptyCols
}
