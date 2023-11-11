package aoc207

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	SolutionPart1() // 386
	SolutionPart2() // 208
}

func TestContainsDuplicates(t *testing.T) {
	data := map[string]bool{
		"aa bb cc dd ee":  true,
		"aa bb cc dd aa":  false,
		"aa bb cc dd aaa": true,
	}

	for k, v := range data {
		if ContainsDuplicates(k) == v {
			fmt.Printf("For %s it works!\n", k)
		} else {
			t.Errorf("For %s it does not work!\n", k)
		}
	}
}

func TestContainsAnagrams(t *testing.T) {
	data := map[string]bool{
		"abcde fghij":              false,
		"abcde xyz ecdab":          true,
		"a ab abc abd abf abj":     false,
		"iiii oiii ooii oooi oooo": false,
		"oiii ioii iioi iiio":      true,
		"bcl nemims udwkmlm nokck tkwny ulkihcu knwty pngamqg yxtphkn kuihlcu": true,
	}

	for k, v := range data {
		if ContainsAnagrams(k) == v {
			fmt.Printf("For %s it works!\n", k)
		} else {
			t.Errorf("For %s it does not work!\n", k)
		}
	}
}
