package aoc2023

import (
	"fmt"
	"strconv"
	"strings"
)

var CACHE = map[string]int{}

func CountPossibleArrangement(input []string, factor int) int {
	result := 0
	for _, line := range input {
		rowInfo := strings.Split(line, " ")
		springs := UnfoldRecord(rowInfo[0], factor, "?")
		groupStr := strings.Split(UnfoldRecord(rowInfo[1], factor, ","), ",")
		groups := []int{}
		for _, x := range groupStr {
			group, _ := strconv.Atoi(x)
			groups = append(groups, group)
		}
		count := Permutation(springs, groups)

		result += count
	}
	return result
}

func Permutation(input string, groups []int) int {
	if len(groups) == 0 && len(input) == 0 {
		return 1
	} else if len(input) == 0 {
		return 0
	} else if len(groups)*2-1 > len(input) {
		return 0
	}
	key := input + "|" + fmt.Sprint(groups)
	if val, ok := CACHE[key]; ok {
		return val
	}

	permutations := 0
	if input[0] == '.' {
		permutations = Permutation(input[1:], groups)
	} else if input[0] == '?' {
		permutations = Permutation("."+input[1:], groups) + Permutation("#"+input[1:], groups)
	} else {
		if len(groups) > 0 {
			group := groups[0]
			if group <= len(input) {
				subLine := input[:group]
				if strings.Contains(subLine, ".") {
					return 0
				} else {
					if len(input) == len(subLine) && len(groups[1:]) == 0 {
						permutations = 1
					} else if len(subLine) < len(input) && input[group] == '.' {
						permutations = Permutation(input[group+1:], groups[1:])
					} else if len(subLine) < len(input) && input[group] == '?' {
						permutations = Permutation("."+input[group+1:], groups[1:])
					}
				}
			}
		}
	}
	CACHE[key] = permutations
	return permutations
}

func UnfoldRecord(line string, factor int, joiner string) string {
	if factor == 1 {
		return line
	}
	result := line
	for i := 1; i < factor; i++ {
		result += joiner + line
	}
	return result
}
