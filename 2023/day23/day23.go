package aoc2023

import (
	"strings"
)

var DIRS = map[byte][2]int{
	'^': {0, -1}, // up
	'<': {-1, 0}, // left
	'v': {0, 1},  // down
	'>': {1, 0},  // right
}

type Point struct {
	x int
	y int
}

type Node struct {
	x       int
	y       int
	visited []Point
}

func (n *Node) isVisited(newNodePos Point) bool {
	for _, v := range n.visited {
		if v == newNodePos {
			return true
		}
	}
	return false
}

func FindTheLongestDistanceWithoutSlopes(input []string) int {
	points := []Point{}
	startPos := Point{x: strings.Index(input[0], "."), y: 0}
	endPos := Point{x: strings.Index(input[len(input)-1], "."), y: len(input) - 1}
	points = append(points, startPos)
	points = append(points, endPos)

	// find conjunctions of paths
	for y, line := range input {
		line = strings.ReplaceAll(line, "<", ".")
		line = strings.ReplaceAll(line, ">", ".")
		line = strings.ReplaceAll(line, "^", ".")
		line = strings.ReplaceAll(line, "v", ".")
		input[y] = line
		for x, field := range line {
			if field == '#' {
				continue
			}
			neighbours := Neighbours(input, x, y)
			if len(neighbours) > 2 {
				points = append(points, Point{x, y})
			}
		}
	}

	// create a list/matrix of distances between conjunctions
	distances := CreateDistanceMatrix(points, input)
	return LongestDistance(distances, startPos, endPos)
}

func CreateDistanceMatrix(points []Point, input []string) map[Point]map[Point]int {
	distances := map[Point]map[Point]int{}
	for i := 0; i < len(points)-1; i++ {
		p1 := points[i]
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]
			dist := FindTheLongestPathBetweenPoints(input, p1, p2, points)
			if dist > 0 {
				if distances[p1] == nil {
					distances[p1] = map[Point]int{}
				}
				if distances[p2] == nil {
					distances[p2] = map[Point]int{}
				}
				distances[p1][p2] = dist
				distances[p2][p1] = dist
			}
		}
	}
	return distances
}

// DFS to find the longest distance between conjunctions
func LongestDistance(distances map[Point]map[Point]int, start, end Point) int {
	visited := make(map[Point]bool)

	var dfs func(current, end Point) int
	dfs = func(current, end Point) int {
		if current == end {
			return 0
		}

		visited[current] = true
		longestDistance := -1

		for point, distance := range distances[current] {
			if !visited[point] {
				newDistance := dfs(point, end)
				if newDistance != -1 && newDistance+distance > longestDistance {
					longestDistance = newDistance + distance
				}
			}
		}

		visited[current] = false
		return longestDistance
	}

	return dfs(start, end)
}

func FindTheLongestPath(input []string) int {
	startPos := Point{strings.Index(input[0], "."), 0}
	endPos := Point{strings.Index(input[len(input)-1], "."), len(input) - 1}
	return FindTheLongestPathBetweenPoints(input, startPos, endPos, []Point{})
}

// for part2 breakPoints are points that are conjunctions of paths. If the algorithm encounters these
// points should skip paths with these points
func FindTheLongestPathBetweenPoints(input []string, startPos, endPos Point, breakPoints []Point) int {
	start := Node{x: startPos.x, y: startPos.y}
	maxCost := map[Point]int{startPos: 0}

	queue := []Node{start}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if len(breakPoints) > 0 && ShouldSkip(breakPoints, current, startPos) {
			continue
		}

		nextSteps := FindNextSteps(input, current)
		for _, next := range nextSteps {
			if current.isVisited(next) || next == startPos {
				continue
			}

			currCost := maxCost[next]
			cur := maxCost[Point{current.x, current.y}]
			if currCost == 0 || currCost < cur+1 {
				maxCost[next] = cur + 1
				path := current.visited
				copy := append(make([]Point, 0, len(path)), path...)
				copy = append(copy, next)
				queue = append(queue, Node{x: next.x, y: next.y, visited: copy})
			}
		}

	}
	return maxCost[endPos]
}

func ShouldSkip(breakPoints []Point, current Node, startPos Point) bool {
	for _, s := range breakPoints {
		if current.x == s.x && current.y == s.y && s != startPos {
			return true
		}
	}
	return false
}

func FindNextSteps(input []string, cur Node) []Point {
	if input[cur.y][cur.x] == '.' {
		return Neighbours(input, cur.x, cur.y)
	} else {
		forceDir := DIRS[input[cur.y][cur.x]]
		return []Point{{cur.x + forceDir[0], cur.y + forceDir[1]}}
	}
}

func Neighbours(input []string, x, y int) []Point {
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	paths := []Point{}
	for _, dir := range directions {
		maxX := len(input[0])
		maxY := len(input)
		newX := x + dir[0]
		newY := y + dir[1]
		if newX >= 0 && newY >= 0 && newX < maxX && newY < maxY && input[newY][newX] != '#' {
			paths = append(paths, Point{x: newX, y: newY})
		}
	}
	return paths
}
