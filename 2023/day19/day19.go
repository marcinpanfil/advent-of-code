package aoc2023

import (
	"strconv"
	"strings"
)

var MIN = 1
var MAX = 4000

type Rating struct {
	x, m, a, s int
}

func (r *Rating) setValue(rating rune, value int) {
	switch rating {
	case 'x':
		r.x = value
	case 'm':
		r.m = value
	case 'a':
		r.a = value
	case 's':
		r.s = value
	default:
		panic("wrong rating!")
	}
}

func (r Rating) getValue(rating rune) int {
	switch rating {
	case 'x':
		return r.x
	case 'm':
		return r.m
	case 'a':
		return r.a
	case 's':
		return r.s
	}
	panic("wrong rating!")
}

func (r *Rating) Copy() Rating {
	return Rating{x: r.x, m: r.m, a: r.a, s: r.s}
}

type Workflow struct {
	name  string
	rules []string
}

func CalculateRating(input []string) int {
	workflows, ratings := ParseInput(input)
	result := 0
	for _, rating := range ratings {
		result += ProcessRating(rating, workflows)
	}
	return result
}

func CalulcatePossibleCombinations(input []string) int {
	workflows, _ := ParseInput(input)
	minR := Rating{x: MIN, m: MIN, a: MIN, s: MIN}
	maxR := Rating{x: MAX, m: MAX, a: MAX, s: MAX}
	return ProcessRange(workflows, "in", minR, maxR)
}

func ProcessRange(workflows map[string]Workflow, curr string, minR, maxR Rating) int {
	currWorkflow := workflows[curr]
	rules := currWorkflow.rules
	result := 0

	for _, rule := range rules {
		if rule == "A" {
			result += Calculate(minR, maxR)
		} else if !strings.Contains(rule, ":") && rule != "R" {
			result += ProcessRange(workflows, rule, minR, maxR)
		} else if rule != "R" {
			dst, part, comparison, value := ParseRuleString(rule)
			if comparison == '>' {
				if maxR.getValue(part) > value {
					newMax := maxR.Copy()
					newMin := minR.Copy()
					newMin.setValue(part, max(newMin.getValue(part), value+1))
					result += ProcessNextStep(workflows, dst, newMin, newMax)
				}
				maxR.setValue(part, value)
			} else if comparison == '<' {
				if minR.getValue(part) < value {
					newMax := maxR.Copy()
					newMin := minR.Copy()
					newMax.setValue(part, min(newMax.getValue(part), value-1))
					result += ProcessNextStep(workflows, dst, newMin, newMax)
				}
				minR.setValue(part, value)
			} else {
				panic("Wrong comparison!")
			}
		}
	}
	return result
}

func ProcessNextStep(workflows map[string]Workflow, dst string, min, max Rating) int {
	if dst == "A" {
		return Calculate(min, max)
	} else if dst != "R" {
		return ProcessRange(workflows, dst, min, max)
	}
	return 0
}

func ParseRuleString(rule string) (string, rune, byte, int) {
	ruleData := strings.Split(rule, ":")
	condition := ruleData[0]
	destination := ruleData[1]
	part := rune(condition[0])
	comparison := condition[1]
	value, _ := strconv.Atoi(condition[2:])
	return destination, part, comparison, value
}

func Calculate(min Rating, max Rating) int {
	return (1 + max.x - min.x) * (1 + max.m - min.m) *
		(1 + max.a - min.a) * (1 + max.s - min.s)
}

func ProcessRating(rating Rating, workflows map[string]Workflow) int {
	workflow := workflows["in"]
	for {
		for _, rule := range workflow.rules {
			result := ProcessRule(rating, rule)
			if result == "A" {
				return rating.x + rating.m + rating.a + rating.s
			} else if result == "R" {
				return 0
			} else if result != "" {
				workflow = workflows[result]
				break
			}
		}
	}
}

func ProcessRule(rating Rating, rule string) string {
	if rule == "A" || rule == "R" || !strings.Contains(rule, ":") {
		return rule
	}
	destination, part, comparison, value := ParseRuleString(rule)
	if comparison == '>' {
		if rating.getValue(rune(part)) > value {
			return destination
		}
	} else if comparison == '<' {
		if rating.getValue(rune(part)) < value {
			return destination
		}
	}
	return ""
}

func ParseInput(input []string) (map[string]Workflow, []Rating) {
	workflowPart := true
	workflows := map[string]Workflow{}
	ratings := []Rating{}
	for _, line := range input {
		if line == "" {
			workflowPart = false
			continue
		}
		if workflowPart {
			openIdx := strings.Index(line, "{")
			name := line[:openIdx]
			rulesStr := line[openIdx+1 : len(line)-1]
			rules := strings.Split(rulesStr, ",")
			workflows[name] = Workflow{name: name, rules: rules}
		} else {
			params := strings.Split(line[1:len(line)-1], ",")
			mapping := map[rune]int{}
			for _, param := range params {
				mapping[rune(param[0])], _ = strconv.Atoi(param[2:])
			}
			ratings = append(ratings, Rating{x: mapping['x'], m: mapping['m'], a: mapping['a'], s: mapping['s']})
		}
	}
	return workflows, ratings
}
