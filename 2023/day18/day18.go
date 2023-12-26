package aoc2023

import (
	"advent-of-code/util/math"
	"regexp"
	"strconv"
	"strings"
)

var DIRS = map[string][2]int{
	"U": {0, 1},
	"D": {0, -1},
	"R": {1, 0},
	"L": {-1, 0},
}

type Instruction struct {
	dir   string
	steps int
	color string
}

var REG, _ = regexp.Compile("[#]+")

func DigInterior(input []string) int {
	instructions := ParseInput(input)
	currX := 0
	currY := 0
	points := [][2]int{{currX, currY}}
	trenches := 0
	for _, instruction := range instructions {
		dir := DIRS[instruction.dir]
		currX += dir[0] * instruction.steps
		currY += dir[1] * instruction.steps
		points = append(points, [2]int{currX, currY})
		trenches += instruction.steps
	}

	return PointsInside(points) + trenches
}

func DigInteriorWithHex(input []string) int {
	instructions := ParseInput(input)
	currX := 0
	currY := 0
	points := [][2]int{{currX, currY}}
	trenches := 0
	for _, instruction := range instructions {
		steps, _ := strconv.ParseInt(instruction.color[:5], 16, 64)
		dirSymbol := string(instruction.color[5])
		dir := GetDirection(dirSymbol)

		currX += dir[0] * int(steps)
		currY += dir[1] * int(steps)
		points = append(points, [2]int{currX, currY})
		trenches += int(steps)
	}

	return PointsInside(points) + trenches
}

func GetDirection(dirSymbol string) [2]int {
	if dirSymbol == "0" {
		return DIRS["R"]
	} else if dirSymbol == "1" {
		return DIRS["D"]
	} else if dirSymbol == "2" {
		return DIRS["L"]
	} else if dirSymbol == "3" {
		return DIRS["U"]
	} else {
		panic("Wrong!")
	}
}

func PointsInside(points [][2]int) int {
	area := 0
	boundary := 0
	for i := 0; i < len(points); i++ {
		p1 := points[i]
		p2 := points[(i+1)%len(points)]

		area += ((p1[0] * p2[1]) - (p1[1] * p2[0]))
		boundary += math.Abs(p1[0] - p2[0] + p1[1] - p2[1])
	}
	return math.Abs(area/2) - (boundary / 2) + 1
}

func ParseInput(input []string) []Instruction {
	instructions := []Instruction{}
	for _, l := range input {
		data := strings.Split(l, " ")
		steps, _ := strconv.Atoi(data[1])
		instructions = append(instructions, Instruction{dir: data[0], steps: steps, color: data[2][2 : len(data[2])-1]})
	}
	return instructions
}
