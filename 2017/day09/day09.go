package aoc2017

var GROUP_START = "{"
var GROUP_END = "}"
var GARBAGE_START = "<"
var GARBAGE_END = ">"
var CANCEL_CHAR = "!"

func CountCharsWithinGarbage(input string) int {
	isGarbage := false
	cancledCount := 0
	isCancled := false
	for _, char := range input {
		if string(char) == CANCEL_CHAR && !isCancled {
			isCancled = true
			continue
		}
		if !isGarbage && string(char) == GARBAGE_START && !isCancled {
			isGarbage = true
			continue
		}

		if isGarbage && string(char) == GARBAGE_END && !isCancled {
			isGarbage = false
			continue
		}
		if isGarbage && !isCancled && string(char) != CANCEL_CHAR {
			cancledCount++
		}

		isCancled = false
	}
	return cancledCount
}

func FindTotalScore(input string) int {
	isGarbage := false
	groupCounter := 0
	totalScore := 0
	currScore := 0
	isCancled := false
	for _, char := range input {
		if string(char) == CANCEL_CHAR && !isCancled {
			isCancled = true
			continue
		}
		if !isGarbage && string(char) == GARBAGE_START && !isCancled {
			isGarbage = true
		}

		if isGarbage && string(char) == GARBAGE_END && !isCancled {
			isGarbage = false
		}

		if !isGarbage && string(char) == GROUP_START && !isCancled {
			groupCounter += 1
			currScore += 1
		}

		if !isGarbage && string(char) == GROUP_END && !isCancled {
			groupCounter -= 1
			totalScore += currScore
			currScore -= 1
			if groupCounter == 0 {
				currScore = 0
			}
		}

		isCancled = false
	}
	return totalScore
}
