package aoc2023

import "math"

func FindMinPath(input []string, maxSteps int, minSteps int) int {
	graph := ParseToInts(input)
	minDistance := FindPathWithMinHeat(graph, maxSteps, minSteps)
	return minDistance
}

type Node struct {
	x     int
	y     int
	steps int
	dir   [2]int
}

var DIRS = [][2]int{
	{1, 0},
	{-1, 0},
	{0, -1},
	{0, 1},
}

func FindPathWithMinHeat(graph [][]int, maxSteps int, minSteps int) int {
	maxX := len(graph[0]) - 1
	maxY := len(graph) - 1

	queue := []Node{}
	queue = append(queue, Node{x: 0, y: 0, steps: 0, dir: [2]int{1, 0}})
	queue = append(queue, Node{x: 0, y: 0, steps: 0, dir: [2]int{0, 1}})

	visited := map[Node]int{
		{x: 0, y: 0, steps: 1, dir: [2]int{1, 0}}: graph[0][0],
		{x: 0, y: 0, steps: 1, dir: [2]int{0, 1}}: graph[0][0],
	}

	bestResult := math.MaxInt
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.x == maxX && current.y == maxY && current.steps >= minSteps {
			bestResult = min(visited[current], bestResult)
		}

		dirs := GetValidDirs(current.dir, [2]int{current.x, current.y}, maxX, maxY)
		for _, dir := range dirs {
			newX := current.x + dir[0]
			newY := current.y + dir[1]
			newDist := visited[current] + graph[newY][newX]
			steps := 0
			if (current.dir == dir && current.steps+1 > maxSteps) || (current.dir != dir && current.steps < minSteps) {
				continue
			} else if current.dir == dir {
				steps = current.steps + 1
			} else {
				steps = 1
			}
			newNode := Node{x: newX, y: newY, steps: steps, dir: dir}

			if dist, isVisited := visited[newNode]; !isVisited || dist > newDist {
				queue = append(queue, newNode)
				visited[newNode] = newDist
			}
		}

	}
	return bestResult
}

func GetValidDirs(cur [2]int, pos [2]int, maxX, maxY int) [][2]int {
	possibleDirs := [][2]int{}
	for _, dir := range DIRS {
		// can't go in opposite direction (can't "go back")
		if (cur[0] == -dir[0] && cur[1] == 0) || (cur[1] == -dir[1] && cur[0] == 0) {
			continue
		}

		newX := pos[0] + dir[0]
		newY := pos[1] + dir[1]

		if newX >= 0 && newY >= 0 && newX <= maxX && newY <= maxY {
			possibleDirs = append(possibleDirs, dir)
		}

	}
	return possibleDirs
}

func ParseToInts(input []string) [][]int {
	result := [][]int{}
	for _, line := range input {
		ints := []int{}
		for _, numAsStr := range line {
			num := int(numAsStr - '0')
			ints = append(ints, num)
		}
		result = append(result, ints)
	}
	return result
}
