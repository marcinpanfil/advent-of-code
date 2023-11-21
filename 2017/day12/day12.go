package aoc2017

import (
	"advent-of-code/util/array"
	"slices"
	"strconv"
	"strings"
)

func CountProgramsConnectedTo0(input []string) int {
	data := ParseInputData(input)

	totalCount := 0
	for root, programs := range data {
		if slices.Contains(programs, 0) || root == 0 {
			totalCount++
		} else {
			len := FindPathBetweenPrograms(data, root, 0)
			if len > 0 {
				totalCount++
				continue
			}
		}
	}
	return totalCount
}

func CountGroups(input []string) int {
	data := ParseInputData(input)
	totalCount := 0
	programKeys := []int{}
	for k := range data {
		programKeys = array.AppendIfNotPresent(programKeys, k)
	}

	alreadyInGroup := []int{}
	for root, programs := range data {
		if slices.Contains(alreadyInGroup, root) {
			continue
		}
		for _, programKey := range programKeys {
			if slices.Contains(programs, programKey) || root == programKey {
				alreadyInGroup = array.AppendIfNotPresent(alreadyInGroup, programKey)
			} else {
				len := FindPathBetweenPrograms(data, root, programKey)
				if len > 0 {
					alreadyInGroup = array.AppendIfNotPresent(alreadyInGroup, programKey)
				}
			}
		}
		totalCount++
	}

	return totalCount
}

// some kind of BFS algorithm, if length > 0 - there's a path
func FindPathBetweenPrograms(data map[int][]int, start, end int) int {
	directNeighbors := getNeighbors(data, start)
	if len(directNeighbors) == 0 {
		return 0
	}
	visited := map[int]bool{start: true}
	queue := []int{start}
	length := 0
	for len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]

			if current == end {
				return length
			}

			for _, neighbor := range getNeighbors(data, current) {
				if visited[neighbor] {
					continue
				}
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
		length++
	}

	// didn't find any path
	return 0
}

func getNeighbors(data map[int][]int, program int) []int {
	neighbors := data[program]
	index := slices.Index(neighbors, program)
	if index >= 0 {
		return append(neighbors[:index], neighbors[index+1:]...)
	} else {
		return neighbors
	}
}

func ParseInputData(input []string) map[int][]int {
	data := map[int][]int{}

	for _, line := range input {
		d2d := strings.Split(line, " <-> ")
		programsAsStr := strings.Split(d2d[1], ", ")
		programs := []int{}
		for _, program_as_str := range programsAsStr {
			program, _ := strconv.Atoi(program_as_str)
			programs = array.AppendIfNotPresent(programs, program)
		}
		rootProgram, _ := strconv.Atoi(d2d[0])
		data[rootProgram] = programs
	}
	return data
}
