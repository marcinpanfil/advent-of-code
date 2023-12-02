package aoc2023

import (
	"math"
	"strconv"
	"strings"
)

type Game struct {
	id   int
	sets []map[string]int
}

func SumOfPossibleGames(input []string) int {
	games := ParseLines(input)
	maxCount := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	result := 0
	for _, game := range games {
		hasMore := false
		for _, cubeSet := range game.sets {
			if !hasMore {
				for color, count := range cubeSet {
					if count > maxCount[color] {
						hasMore = true
						break
					}
				}
			}
		}
		if !hasMore {
			result += game.id
		}
	}
	return result
}

func FindMinNumberOfCubes(input []string) int {
	games := ParseLines(input)

	result := 0
	for _, game := range games {
		maxCount := map[string]int{
			"red":   math.MinInt32,
			"green": math.MinInt32,
			"blue":  math.MinInt32,
		}
		for _, cubeSet := range game.sets {
			for color, count := range cubeSet {
				if maxCount[color] < count {
					maxCount[color] = count
				}
			}
		}
		result += maxCount["red"] * maxCount["green"] * maxCount["blue"]
	}
	return result

}

func ParseLines(input []string) []Game {
	games := []Game{}
	for id, line := range input {
		skipPartIdx := strings.Index(line, ": ") + 2
		line = line[skipPartIdx:]
		setsStr := strings.Split(line, "; ")
		sets := []map[string]int{}
		for _, setStr := range setsStr {
			cubesStr := strings.Split(setStr, ", ")
			cubes := map[string]int{}
			for _, cubeStr := range cubesStr {
				data := strings.Split(cubeStr, " ")
				count, _ := strconv.Atoi(data[0])
				color := data[1]
				cubes[color] = count
			}
			sets = append(sets, cubes)
		}
		game := Game{id: id + 1, sets: sets}
		games = append(games, game)
	}

	return games
}
