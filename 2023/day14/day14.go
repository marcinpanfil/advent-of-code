package aoc2023

import (
	"advent-of-code/util/array"
	"strings"
)

func CalculateTotalLoad(input []string) int {
	TiltRocks(input)
	return CalculateLoad(input)
}

func CalculateTotalLoadAfter1T(input []string) int {
	results := map[string]int{}
	for i := 0; i < 1000000000; i++ {
		key := strings.Join(input, " ")
		if results[key] == 0 {
			results[key] = i
		} else if (1_000_000_000-i)%(results[key]-i) == 0 {
			return CalculateLoad(input)
		}
		for r := 0; r < 4; r++ {
			TiltRocks(input)
			input = RotateArray90Degrees(input)
		}
	}
	return -1
}

func TiltRocks(input []string) {
	for i := 1; i < len(input); i++ {
		line := input[i]
		moveableRockIds := array.IndexesOf[rune]([]rune(line), 'O')
		for _, id := range moveableRockIds {
			newLvl := FindTheLowestPos(input, i, id)
			input[newLvl] = input[newLvl][:id] + "O" + input[newLvl][id+1:]
			if newLvl < i {
				input[i] = input[i][:id] + "." + input[i][id+1:]
			}
		}
	}
}

func RotateArray90Degrees(input []string) []string {
	rows := len(input)
	cols := len(input[0])
	rotated := make([]string, cols)

	for i := 0; i < rows; i++ {
		rotated[i] = strings.Repeat(" ", cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rotated[j] = rotated[j][:rows-i-1] + string(input[i][j]) + rotated[j][rows-i:]
		}
	}

	return rotated
}

func CalculateLoad(input []string) int {
	loopRes := 0
	for i, line := range input {
		moveableRockIds := array.IndexesOf[rune]([]rune(line), 'O')
		loopRes += len(moveableRockIds) * (len(line) - i)
	}
	return loopRes
}

func FindTheLowestPos(input []string, curLvl int, id int) int {
	for i := curLvl - 1; i >= 0; i-- {
		if input[i][id] == '#' || input[i][id] == 'O' {
			return i + 1
		}
	}
	return 0
}
