package aoc2023

import (
	"strings"
)

func FindReflectsWithSmudge(mirrors []string, smudgeCount int) int {
	result := 0
	for _, mirror := range mirrors {
		lines := strings.Split(mirror, "\n")
		result += CalculateReflection(lines, 100, smudgeCount)
		colLines := GetColumnMirror(lines)
		result += CalculateReflection(colLines, 1, smudgeCount)
	}
	return result
}

func CalculateReflection(lines []string, factor int, smudgeCount int) int {
	result := 0
	for i := 0; i < len(lines)-1; i++ {
		incorrect := 0
		for j := 0; j < len(lines); j++ {
			curr := i - j
			next := i + 1 + j
			if curr < len(lines) && next < len(lines) && curr >= 0 {
				incorrect += len(FindStringDifferences(lines[curr], lines[next]))
			}
		}
		if incorrect == smudgeCount {
			result += factor * (i + 1)
		}
	}
	return result
}

func GetColumnMirror(lines []string) []string {
	colLines := make([]string, len(lines[0]))
	for _, line := range lines {
		for j, c := range line {
			colLines[j] += string(c)
		}
	}
	return colLines
}

func FindStringDifferences(first string, second string) []int {
	differences := []int{}
	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			differences = append(differences, i)
		}
	}
	return differences
}
