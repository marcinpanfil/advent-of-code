package aoc2017

import (
	"advent-of-code/util/array"
	"fmt"
	"slices"
)

func KnotTyingHash(lenghts []int, maxRange int) int {
	numbers := array.GenerateSequence(0, maxRange-1)
	skipSize := 0
	curPos := 0
	for _, l := range lenghts {
		ReverseList(l, numbers, curPos)
		curPos = (curPos + skipSize + l) % len(numbers)
		skipSize++
	}
	return numbers[0] * numbers[1]
}

func CalculateKnotHash(input string) string {
	toAscii := []byte(input)
	toAscii = append(toAscii, 17, 31, 73, 47, 23)

	numbers := array.GenerateSequence(0, 255)
	skipSize := 0
	curPos := 0
	for rund := 0; rund < 64; rund++ {
		for _, l := range toAscii {
			ReverseList(int(l), numbers, curPos)
			curPos = (curPos + skipSize + int(l)) % len(numbers)
			skipSize++
		}
	}
	dashHash := CalculateDashHash(numbers)
	return ParseHashAsHex(dashHash)
}

func ReverseList(l int, numbers []int, curPos int) {
	subList := make([]int, l)
	CopyWithWrap(subList, numbers, curPos)
	slices.Reverse(subList)
	PasteWithWrap(numbers, subList, curPos)
}

func ParseHashAsHex(dashHash [16]int) string {
	result := ""
	for _, hash := range dashHash {
		result += fmt.Sprintf("%02x", hash)
	}
	return result
}

func CalculateDashHash(numbers []int) [16]int {
	var dashHash [16]int
	for block := 0; block < 16; block++ {
		var xorVar int
		for i := 0; i < 16; i++ {
			xorVar ^= numbers[block*16+i]
		}
		dashHash[block] = xorVar
	}
	return dashHash
}

func CopyWithWrap(dst []int, src []int, start int) {
	for i := range dst {
		dst[i] = src[(i+start)%len(src)]
	}
}

func PasteWithWrap(dst []int, src []int, start int) {
	for i := range src {
		dst[(i+start)%len(dst)] = src[i]
	}
}
