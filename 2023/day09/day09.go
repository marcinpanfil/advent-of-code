package aoc2023

import (
	"strconv"
	"strings"
)

func ExtrapolateValues(input []string) (int, int) {
	result1, result2 := 0, 0
	for _, line := range input {
		valuesAsStr := strings.Split(line, " ")
		values := []int{}
		for _, valueAsStr := range valuesAsStr {
			value, _ := strconv.Atoi(valueAsStr)
			values = append(values, value)
		}
		extraMap := map[int][]int{0: values}
		ExtrapolateNextLvls(extraMap, 0)
		r1, r2 := CalculateFirstLastValues(extraMap)
		result1 += r1
		result2 += r2
	}
	return result1, result2
}

func ExtrapolateNextLvls(values map[int][]int, lvl int) {
	newValues := make([]int, len(values[lvl])-1)
	allZeros := true
	for i := 0; i < len(newValues); i++ {
		value1 := values[lvl][i]
		value2 := values[lvl][i+1]
		newValues[i] = value2 - value1
		if newValues[i] != 0 && allZeros {
			allZeros = false
		}
	}
	values[lvl+1] = newValues
	if !allZeros {
		ExtrapolateNextLvls(values, lvl+1)
	}
}

func CalculateFirstLastValues(values map[int][]int) (int, int) {
	diffLast := 0
	diffFirst := 0
	for i := len(values) - 2; i >= 0; i-- {
		lowerExtrpolaration := values[i+1]
		lastExtrpolaration := values[i]
		diffLast = lastExtrpolaration[len(lastExtrpolaration)-1] + lowerExtrpolaration[len(lowerExtrpolaration)-1]
		diffFirst = lastExtrpolaration[0] - lowerExtrpolaration[0]
		lastExtrpolaration = append(lastExtrpolaration, diffLast)
		lastExtrpolaration = append([]int{diffFirst}, lastExtrpolaration...)
		values[i] = lastExtrpolaration
	}
	return values[0][len(values[0])-1], values[0][0]
}
