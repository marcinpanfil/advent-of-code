package aoc2017

import (
	"strconv"
	"strings"
)

type Program struct {
	name        string
	weight      int
	children    []string
	totalWeight int
}

func LoadProgramsData(programsData []string) (map[string]Program, map[string]int) {
	programs := map[string]Program{}
	// stores all non-root programs (without the bottom program)
	childPrograms := map[string]int{}

	replacer := strings.NewReplacer("(", "", ")", "")
	for _, programEntry := range programsData {
		programData := strings.Split(programEntry, " ")
		programName := programData[0]
		programWeight, _ := strconv.Atoi(replacer.Replace(programData[1]))
		children := []string{}
		if len(programData) > 2 {
			for i := 3; i < len(programData); i++ {
				childName := strings.Replace(programData[i], ",", "", 1)
				children = append(children, childName)
				childPrograms[childName] = 1
			}
		}

		programs[programName] = Program{name: programName, weight: programWeight, children: children}
	}
	return programs, childPrograms
}

func FindBottomProgram(programsData []string) string {
	programs, childPrograms := LoadProgramsData(programsData)
	return GetBottomProgramName(programs, childPrograms)
}

func GetBottomProgramName(programs map[string]Program, childPrograms map[string]int) string {
	for programName := range programs {
		if childPrograms[programName] == 0 {
			return programName
		}
	}
	return ""
}

func FindCorrectWeightOfUnballancedProgram(programsData []string) int {
	programs, childPrograms := LoadProgramsData(programsData)
	bottomProgramName := GetBottomProgramName(programs, childPrograms)
	bottomProgram := programs[bottomProgramName]
	bottomProgram.totalWeight = CalculateWeights(bottomProgram, programs)
	return FindCorrectWeight(bottomProgram, programs)
}

func CalculateWeights(parent Program, programs map[string]Program) int {
	total := 0
	for _, childName := range parent.children {
		childProgram := programs[childName]
		weight := childProgram.weight + CalculateWeights(childProgram, programs)
		childProgram.totalWeight = weight
		programs[childName] = childProgram
		total += weight
	}
	return total
}

func FindCorrectWeight(parent Program, programs map[string]Program) int {
	childrenWeights := map[int][]string{}
	childrenBalances := map[string]int{}

	for _, childName := range parent.children {
		childProgram := programs[childName]
		childrenBalance := FindCorrectWeight(childProgram, programs)
		if childrenBalance != 0 {
			return childrenBalance
		}
		childrenBalances[childName] = childrenBalance
		childNames, hasValue := childrenWeights[childProgram.totalWeight]
		if hasValue {
			childrenWeights[childProgram.totalWeight] = append(childNames, childName)
		} else {
			childrenWeights[childProgram.totalWeight] = []string{childName}
		}
	}

	// there is discrepancy between child programs
	var problematicProgram Program = Program{}
	if len(childrenWeights) > 1 {
		for _, childNames := range childrenWeights {
			// if child programs are balanced it means childNames[0] is the root cause of the problem
			if len(childNames) == 1 && childrenBalances[childNames[0]] == 0 {
				problematicProgram = programs[childNames[0]]
			}
		}
	}

	if problematicProgram.name != "" {
		return CalculateCorrectWight(parent, problematicProgram, programs)
	}

	return 0
}

func CalculateCorrectWight(parent Program, problematicProgram Program, programs map[string]Program) int {
	correctWeight := 0
	for _, childName := range parent.children {
		if childName != problematicProgram.name {
			correctWeight = programs[childName].totalWeight
			break
		}
	}
	return problematicProgram.weight + (correctWeight - problematicProgram.totalWeight)
}
