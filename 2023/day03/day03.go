package aoc2023

import (
	"advent-of-code/util/array"
	"fmt"
	"slices"
	"strconv"
)

var DIRS = [][2]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{1, -1},
	{1, 0},
	{1, 1},
	{0, 1},
	{0, -1},
}

func CalcEngine(input []string) int {
	result := 0
	for i, line := range input {
		hasAdj := false
		valueAsStr := ""
		for j, char := range line {
			_, err := strconv.Atoi(string(char))
			if err == nil {
				valueAsStr += string(char)
				if !hasAdj {
					hasAdj = HasAdjacentSymbol(input, line, i, j)
				}
			} else {
				if hasAdj {
					value, _ := strconv.Atoi(valueAsStr)
					result += value
				}
				hasAdj = false
				valueAsStr = ""
			}
		}
		// if the digit is the last char in the line
		if hasAdj {
			value, _ := strconv.Atoi(valueAsStr)
			result += value
		}
	}

	return result
}

func CalcGear(input []string) int {
	result := 0
	asteriskPos := [][2]int{}
	for idx, line := range input {
		asterisksInLine := array.IndexesOf[rune]([]rune(line), '*')
		for _, idInLine := range asterisksInLine {
			asteriskPos = append(asteriskPos, [2]int{idInLine, idx})
		}
	}

	for _, pos := range asteriskPos {
		count := 0
		parts := 1
		visited := [][2]int{}
		for _, dir := range DIRS {
			x := dir[0] + pos[0]
			y := dir[1] + pos[1]
			if x < 0 || y < 0 || x >= len(input[0]) || y >= len(input) {
				continue
			}
			_, err := strconv.Atoi(string(input[y][x]))
			if err == nil && !slices.Contains(visited, [2]int{x, y}) {
				count++
				lineWithNr := input[y]
				valueAsStr := ""
				for k := x; k < len(lineWithNr); k++ {
					val, err := strconv.Atoi(string(lineWithNr[k]))
					if err == nil {
						valueAsStr += fmt.Sprint(val)
						visited = append(visited, [2]int{k, y})
					} else {
						break
					}
				}
				for k := x - 1; k >= 0; k-- {
					val, err := strconv.Atoi(string(lineWithNr[k]))
					if err == nil {
						valueAsStr = fmt.Sprint(val) + valueAsStr
						visited = append(visited, [2]int{k, y})
					} else {
						break
					}
				}
				value, _ := strconv.Atoi(valueAsStr)
				parts *= value
				visited = append(visited, [2]int{x, y})
			}
		}
		if count == 2 {
			result += parts
		}
	}

	return result
}

func HasAdjacentSymbol(input []string, line string, i int, j int) bool {
	for _, dir := range DIRS {
		x := dir[0] + j
		y := dir[1] + i
		if x >= 0 && x < len(line) && y >= 0 && y < len(input) {
			_, err := strconv.Atoi(string(input[y][x]))
			if string(input[y][x]) != "." && err != nil {
				return true
			}
		}
	}
	return false
}
