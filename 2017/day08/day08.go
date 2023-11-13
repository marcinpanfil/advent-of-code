package aoc2017

import (
	maps "advent-of-code/util/map"
	"math"
	"strconv"
	"strings"
)

type Instruction struct {
	register     string
	operation    string
	opValue      int
	condRegister string
	condition    string
	condValue    int
}

func FindTheLargestRegister(instructionsInput []string) (int, int) {
	instructions := []Instruction{}
	registers := map[string]int{}
	for _, instructionInput := range instructionsInput {
		parts := strings.Split(instructionInput, " ")
		opValue, _ := strconv.Atoi(parts[2])
		condValue, _ := strconv.Atoi(parts[6])
		register := parts[0]
		condRegister := parts[4]
		instructions = append(instructions, Instruction{
			register:     register,
			operation:    parts[1],
			opValue:      opValue,
			condRegister: condRegister,
			condition:    parts[5],
			condValue:    condValue,
		})
		registers[register] = 0
	}
	historialMax := Process(instructions, registers)
	return maps.FindMaxInIntValues(registers), historialMax
}

func Process(Instructions []Instruction, registers map[string]int) int {
	historicalMax := math.MinInt64
	for _, instruction := range Instructions {
		if CheckCondition(registers[instruction.condRegister], instruction.condition, instruction.condValue) {
			value := math.MinInt64
			if instruction.operation == "inc" {
				value = registers[instruction.register] + instruction.opValue
				registers[instruction.register] = value
			} else if instruction.operation == "dec" {
				value = registers[instruction.register] - instruction.opValue
				registers[instruction.register] = value
			}
			if value > historicalMax {
				historicalMax = value
			}
		}
	}
	return historicalMax
}

func CheckCondition(regValue int, condition string, value int) bool {
	switch condition {
	case "<":
		return regValue < value
	case ">":
		return regValue > value
	case "==":
		return regValue == value
	case "!=":
		return regValue != value
	case "<=":
		return regValue <= value
	case ">=":
		return regValue >= value
	default:
		panic("Unsupported!")
	}
}
