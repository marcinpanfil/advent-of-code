package aoc2017

import (
	"strconv"
	"strings"
)

func ParseInput(lines []string) map[int]int {
	data := map[int]int{}
	for _, line := range lines {
		lineParsed := strings.Split(line, ": ")
		depth, _ := strconv.Atoi(lineParsed[0])
		_range, _ := strconv.Atoi(lineParsed[1])
		data[depth] = _range
	}
	return data
}

func SecurityScanner(data map[int]int) int {
	result := 0
	for depth, _range := range data {
		// packet will reach pos 0 only in two cases: at start and after (_range-1) * 2 steps
		stepsTo0 := (_range - 1) * 2
		if depth%stepsTo0 == 0 {
			result += (depth * _range)
		}
	}
	return result
}

func FindMaxDelay(data map[int]int) int {
	delay := 0
	for {
		isCaught := false
		for depth, _range := range data {
			stepsTo0 := (_range - 1) * 2
			if (depth+delay)%stepsTo0 == 0 {
				isCaught = true
				break
			}
		}

		if !isCaught {
			return delay
		}
		delay++
	}
}
