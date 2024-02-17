package aoc2017

import (
	"strings"
	"unicode"
)

var DIRS = [][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func GetStart(input []string) [2]int {
	firstLine := input[0]
	posX := strings.LastIndex(firstLine, "|")
	return [2]int{posX, 0}
}

func FindPath(input []string) (string, int) {
	currPos := GetStart(input)
	dir := [2]int{0, 1}
	result := []byte{}
	steps := 0

	for dir[0] != 0 || dir[1] != 0 {
		steps++
		currPos = [2]int{currPos[0] + dir[0], currPos[1] + dir[1]}
		curr := input[currPos[1]][currPos[0]]
		if unicode.IsLetter(rune(curr)) {
			result = append(result, curr)
		} else if curr == '+' {
			dir = FindNextDir(input, currPos, dir)
		}
		if curr == ' ' || (dir[0] == 0 && dir[1] == 0) {
			break
		}
	}
	return string(result), steps
}

func FindNextDir(input []string, currPos [2]int, currDir [2]int) [2]int {
	for _, dir := range DIRS {
		if (dir[0] != 0 && dir[0] == -currDir[0]) || (dir[1] != 0 && dir[1] == -currDir[1]) {
			continue
		}
		newPos := [2]int{currPos[0] + dir[0], currPos[1] + dir[1]}
		if input[newPos[1]][newPos[0]] != ' ' {
			return dir
		}
	}
	return [2]int{0, 0}
}
