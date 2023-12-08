package aoc2023

import (
	"advent-of-code/util/math"
	"strings"
)

type Instruction struct {
	name  string
	left  string
	right string
}

func CalculateSteps(input []string) int {
	count := 0
	steps, instructions := ParseInput(input)
	currIns := instructions["AAA"]
	for {
		for _, step := range steps {
			count++
			if step == 'L' {
				currIns = instructions[currIns.left]
			} else if step == 'R' {
				currIns = instructions[currIns.right]
			}
			if currIns.name == "ZZZ" {
				return count
			}
		}
	}
}

func CalculateStepsSimultaneously(input []string) int {
	count := 0
	steps, instructions := ParseInput(input)
	nodes := FindNodesEndingWithA(instructions)
	stepOfFirstOccurrence := make([]int, len(nodes))
	for {
		for _, step := range steps {
			count++
			IterateViaNodes(nodes, step, instructions)
			UpdateOccurrences(nodes, stepOfFirstOccurrence, count)
			shouldBreak := AllOccurrencesFilled(stepOfFirstOccurrence)
			if shouldBreak {
				return math.LCM(stepOfFirstOccurrence...)
			}
		}
	}

}

func UpdateOccurrences(nodes []string, stepOfFirstOccurrence []int, count int) {
	for i, name := range nodes {
		if strings.LastIndex(name, "Z") == 2 {
			if stepOfFirstOccurrence[i] == 0 {
				stepOfFirstOccurrence[i] = count
			}
		}
	}
}

func IterateViaNodes(nodesWithA []string, step rune, instructions map[string]Instruction) {
	for i, node := range nodesWithA {
		if step == 'L' {
			nodesWithA[i] = instructions[node].left
		} else if step == 'R' {
			nodesWithA[i] = instructions[node].right
		}
	}
}

func AllOccurrencesFilled(firstOccurrence []int) bool {
	for _, occ := range firstOccurrence {
		if occ == 0 {
			return false
		}
	}
	return true
}

func FindNodesEndingWithA(instructions map[string]Instruction) []string {
	nodes := []string{}
	for k := range instructions {
		if strings.LastIndex(k, "A") == 2 {
			nodes = append(nodes, k)
		}
	}
	return nodes
}

func ParseInput(input []string) (string, map[string]Instruction) {
	steps := input[0]
	instructions := map[string]Instruction{}
	for i := 2; i < len(input); i++ {
		line := input[i]
		ins := strings.Split(line, " = ")
		nodes := strings.Split(ins[1][1:9], ", ")
		instructions[ins[0]] = Instruction{name: ins[0], left: nodes[0], right: nodes[1]}
	}
	return steps, instructions
}
