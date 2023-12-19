package aoc2023

import "math"

type Beam struct {
	x   int
	y   int
	dir [2]int
}

func CountFromAllTiles(input []string) int {
	max := math.MinInt
	for i := range input {
		cur := CountAllEnergized(input, -1, i, [2]int{1, 0})
		if cur > max {
			max = cur
		}
		cur = CountAllEnergized(input, len(input), i, [2]int{-1, 0})
		if cur > max {
			max = cur
		}
	}
	for i := range input[0] {
		cur := CountAllEnergized(input, i, -1, [2]int{0, 1})
		if cur > max {
			max = cur
		}
		cur = CountAllEnergized(input, i, len(input), [2]int{0, -1})
		if cur > max {
			max = cur
		}
	}
	return max
}

func CountAllEnergized(input []string, startX, startY int, dir [2]int) int {
	queue := []Beam{}
	visited := map[Beam]bool{}
	start := Beam{x: startX, y: startY, dir: dir}
	queue = append(queue, start)
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		nextX := current.x + current.dir[0]
		nextY := current.y + current.dir[1]
		if visited[current] {
			continue
		}
		visited[current] = true
		if nextX >= 0 && nextX < len(input[0]) && nextY >= 0 && nextY < len(input) {
			if input[nextY][nextX] == '.' {
				queue = append(queue, Beam{x: nextX, y: nextY, dir: current.dir})
			} else if input[nextY][nextX] == '|' {
				if current.dir[1] == 0 {
					for _, i := range [2]int{-1, 1} {
						dir := [2]int{0, i}
						b := Beam{x: nextX, y: nextY, dir: dir}
						queue = append(queue, b)
					}
				} else {
					queue = append(queue, Beam{x: nextX, y: nextY, dir: current.dir})
				}
			} else if input[nextY][nextX] == '-' {
				if current.dir[0] == 0 {
					for _, i := range [2]int{-1, 1} {
						dir := [2]int{i, 0}
						b := Beam{x: nextX, y: nextY, dir: dir}
						queue = append(queue, b)
					}
				} else {
					queue = append(queue, Beam{x: nextX, y: nextY, dir: current.dir})
				}
			} else if input[nextY][nextX] == '/' {
				queue = append(queue, Beam{x: nextX, y: nextY, dir: [2]int{-current.dir[1], -current.dir[0]}})
			} else if input[nextY][nextX] == '\\' {
				queue = append(queue, Beam{x: nextX, y: nextY, dir: [2]int{current.dir[1], current.dir[0]}})
			}
		}
	}

	result := map[[2]int]bool{}
	for k, v := range visited {
		if v {
			result[[2]int{k.x, k.y}] = true
		}
	}
	return len(result) - 1
}
