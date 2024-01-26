package aoc2023

import (
	"strings"
)

var DIRS = [][2]int{
	{0, 1},  // up
	{-1, 0}, // left
	{0, -1}, // down
	{1, 0},  // right
}

/*
During the analysis of input and results for different steps, I assumed:
  - there are no rocks in the same line and column as starting point - it means in the first iteration,
    there are max 65 steps to reach the border of the map, later on, the number of steps equals to the size of the map (131).
  - based on the assumptions above: 26501365 (task requirement) steps means that there are 202300 copies of the map required
    -> (26501365-65) / 131 = 202300

I analysed the first 10 sizes of the map (map size =[1,10] and discovered that the delta of the deltas of results (so the second level
delta) for two adjecent sizes (for example, 1 and 2 or 3 and 4) is equal every time. Based on that, I created a linear function to
calculate next delta of results and the final result.
*/
func InterpolationOfPlots(input []string) int64 {
	res1 := ReachGardenPlots(input, 65)
	res2 := ReachGardenPlots(input, 65+131)
	res3 := ReachGardenPlots(input, 65+(131*2))

	firstDeltaLvl1 := res2 - res1
	secondDeltaLvl1 := res3 - res2
	deltaLvl2 := secondDeltaLvl1 - firstDeltaLvl1

	result := int64(res1)
	for i := 1; i <= 202300; i++ {
		result += int64(deltaLvl2)*int64(i) + int64(65)
	}
	return result
}

func ReachGardenPlots(input []string, maxSteps int) int {
	increasedInput := []string{}
	for i := 0; i < 5; i++ {
		for _, line := range input {
			tmpLine := strings.ReplaceAll(line, "S", ".")
			increasedInput = append(increasedInput, strings.Repeat(tmpLine, 5))
		}
	}

	steps := 0
	start := [2]int{len(increasedInput[0]) / 2, len(increasedInput) / 2}
	queue := map[int]map[[2]int]bool{steps: {start: true}}

	for {
		currents := queue[steps]
		steps++
		for current := range currents {
			for _, dir := range DIRS {
				nextX := current[0] + dir[0]
				nextY := current[1] + dir[1]
				neighbor := [2]int{nextX, nextY}
				if nextX < 0 || nextY < 0 || nextX > len(increasedInput[0])-1 || nextY > len(increasedInput)-1 {
					continue
				}
				if increasedInput[nextY][nextX] != '#' {
					stepsMapping := queue[steps]
					if stepsMapping == nil {
						queue[steps] = map[[2]int]bool{neighbor: true}
					} else {
						queue[steps][neighbor] = true
					}
				}
			}
		}
		if steps == maxSteps {
			return len(queue[maxSteps])
		}
	}
}

func FindStart(input []string) [2]int {
	for y, line := range input {
		x := strings.Index(line, "S")
		if x != -1 {
			return [2]int{x, y}
		}
	}
	return [2]int{}
}
