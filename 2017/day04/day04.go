package aoc207

import (
	"advent-of-code/util/file"
	"fmt"
	"slices"
	"strings"
)

func SolutionPart1() {
	passphrases := file.ReadInput()
	correctCount := 0
	for _, phrase := range passphrases {
		if ContainsDuplicates(phrase) {
			correctCount += 1
		}
	}
	fmt.Printf("Result part1: %v\n", correctCount)
}

func SolutionPart2() {
	passphrases := file.ReadInput()
	correctCount := 0
	for _, phrase := range passphrases {
		if !ContainsAnagrams(phrase) {
			correctCount += 1
		}
	}
	fmt.Printf("Result part2: %v\n", correctCount)
}

func ContainsDuplicates(passphrase string) bool {
	words := strings.Fields(passphrase)
	used := []string{}
	for _, word := range words {
		if slices.Contains(used, word) {
			return false
		} else {
			used = append(used, word)
		}
	}
	return true
}

func ContainsAnagrams(passphrase string) bool {
	words := strings.Fields(passphrase)
	for i := 0; i < len(words)-1; i++ {
		word1 := words[i]
		for j := i + 1; j < len(words); j++ {
			word2 := words[j]
			if len(word1) == len(word2) {
				allChars := ContainsTheSameCharacters(word1, word2)
				if allChars {
					return true
				}
			}
		}
	}
	return false
}

func ContainsTheSameCharacters(word1 string, word2 string) bool {
	allChars := true
	for i := 0; i < len(word1); i++ {
		if !strings.Contains(word2, string(word1[i])) {
			return false
		} else {
			word2 = strings.Replace(word2, string(word1[i]), "", 1)
		}
	}
	return allChars
}
