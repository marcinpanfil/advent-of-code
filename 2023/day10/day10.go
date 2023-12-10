package aoc2023

import (
	"advent-of-code/util/array"
	"regexp"
	"sort"
	"strings"
)

var DIRS = map[rune][4]int{
	//    N, E, S, W
	'|': {1, 0, 1, 0},
	'-': {0, 1, 0, 1},
	'L': {1, 1, 0, 0},
	'J': {1, 0, 0, 1},
	'7': {0, 0, 1, 1},
	'F': {0, 1, 1, 0},
}

var REG_FJ, _ = regexp.Compile("[F]{1}[-]*[J]{1}")
var REG_L7, _ = regexp.Compile("[L]{1}[-]*[7]{1}")

func FindLongestLoop(input []string) int {
	start := FindStart(input)
	length, _ := FindLoop(start, input)
	return length - 1
}

func FindLoop(start [2]int, input []string) (int, map[[2]int]bool) {
	queue := [][2]int{start}
	visited := map[[2]int]bool{start: true}
	length := 0
	for len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]

			for _, neighbor := range GetNeighbors(input, current) {
				if visited[neighbor] {
					continue
				}
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
		length++
	}
	return length, visited
}

func CountInsiders(input []string) int {
	start := FindStart(input)
	ReplaceS(input, start)
	_, visited := FindLoop(start, input)
	ReplacePipesNotInLoop(input, visited)

	result := 0

	for _, line := range input {
		lineCount := 0
		dotIndexes := array.IndexesOf[rune]([]rune(line), '.')
		indexes := GetIndexesOfIntersactions(line)
		if len(indexes) > 0 {
			firstC, lastC := indexes[0], indexes[len(indexes)-1]
			for _, index := range dotIndexes {
				if index >= firstC && index <= lastC {
					r1Count := len(FindAllMatches(line[firstC:index+1], REG_FJ))
					r2Count := len(FindAllMatches(line[firstC:index+1], REG_L7))
					count := len(array.IndexesOf[rune]([]rune(line)[firstC:index+1], '|')) + r1Count + r2Count
					if count%2 == 1 {
						lineCount++
					}
				}
			}
		}
		result += lineCount
	}
	return result
}

func GetIndexesOfIntersactions(line string) []int {
	indexes := array.IndexesOf[rune]([]rune(line), '|')
	r1Indexes := FindAllMatches(line, REG_FJ)
	r2Indexes := FindAllMatches(line, REG_L7)
	indexes = append(indexes, r1Indexes...)
	indexes = append(indexes, r2Indexes...)
	sort.Ints(indexes)
	return indexes
}

func ReplacePipesNotInLoop(input []string, visited map[[2]int]bool) {
	for y, line := range input {
		for x, char := range line {
			isVisited := visited[[2]int{x, y}]
			if char != '.' && !isVisited {
				line = line[:x] + "." + line[x+1:]
			}
		}
		input = append(input[:y], append([]string{line}, input[y+1:]...)...)
	}
}

func ReplaceS(input []string, sPos [2]int) {
	currNeighbors := GetNeighbors(input, sPos)
	for dir := range DIRS {
		tmpInput := make([]string, len(input))
		copy(tmpInput, input)
		tmpInput[sPos[1]] = strings.ReplaceAll(tmpInput[sPos[1]], "S", string(dir))
		tmpNeighbors := GetNeighbors(tmpInput, sPos)
		if array.PositionsEquals(currNeighbors, tmpNeighbors) {
			input[sPos[1]] = strings.ReplaceAll(input[sPos[1]], "S", string(dir))
			return
		}
	}
}

func FindAllMatches(line string, r *regexp.Regexp) []int {
	indexes := []int{}
	matches := r.FindAllIndex([]byte(line), -1)
	for _, m := range matches {
		indexes = append(indexes, m[0])
	}
	return indexes
}

var NBR_DIRS = [][2]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func GetNeighbors(input []string, pos [2]int) [][2]int {
	neighbors := [][2]int{}
	curr := rune(input[pos[1]][pos[0]])
	for _, dir := range NBR_DIRS {
		x := dir[0] + pos[0]
		y := dir[1] + pos[1]
		if x >= 0 && x < len(input[0]) && y >= 0 && y < len((input)) {
			neighbor := input[y][x]
			if neighbor == '.' {
				continue
			}
			if curr == 'S' {
				neighbors = GetNeighborsForStart(dir, neighbor, neighbors, x, y)
			} else if dir[1] == -1 && DIRS[rune(curr)][0] == 1 {
				neighbors = append(neighbors, [2]int{x, y})
			} else if dir[1] == 1 && DIRS[rune(curr)][2] == 1 {
				neighbors = append(neighbors, [2]int{x, y})
			} else if dir[0] == -1 && DIRS[rune(curr)][3] == 1 {
				neighbors = append(neighbors, [2]int{x, y})
			} else if dir[0] == 1 && DIRS[rune(curr)][1] == 1 {
				neighbors = append(neighbors, [2]int{x, y})
			}
		}
	}
	return neighbors
}

func GetNeighborsForStart(dir [2]int, neighbor byte, neighbors [][2]int, x int, y int) [][2]int {
	if dir[1] == 1 && DIRS[rune(neighbor)][0] == 1 {
		neighbors = append(neighbors, [2]int{x, y})
	} else if dir[1] == -1 && DIRS[rune(neighbor)][2] == 1 {
		neighbors = append(neighbors, [2]int{x, y})
	} else if dir[0] == 1 && DIRS[rune(neighbor)][3] == 1 {
		neighbors = append(neighbors, [2]int{x, y})
	} else if dir[0] == -1 && DIRS[rune(neighbor)][1] == 1 {
		neighbors = append(neighbors, [2]int{x, y})
	}
	return neighbors
}

func FindStart(input []string) [2]int {
	for i, line := range input {
		for j, str := range line {
			if str == 'S' {
				return [2]int{j, i}
			}
		}
	}
	panic("Something wrong!")
}
