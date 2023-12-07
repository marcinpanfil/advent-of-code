package aoc2023

import (
	"advent-of-code/util/file"
	"fmt"
	"strings"
	"testing"
)

var INPUT = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func TestSolution(t *testing.T) {
	input := file.ReadInput()
	fmt.Println(CalculateTotalWinnings(input, false)) // 252052080
	fmt.Println(CalculateTotalWinnings(input, true))  // 252898370
}

func TestCalculateTotalWinnings(t *testing.T) {
	result := CalculateTotalWinnings(strings.Split(INPUT, "\n"), false)
	if result == 6440 {
		fmt.Println("Works!")
	} else {
		t.Errorf("Wrong value %v!", result)
	}
}

func TestCalculateTotalWinningsWithJ(t *testing.T) {
	result := CalculateTotalWinnings(strings.Split(INPUT, "\n"), true)
	if result == 5905 {
		fmt.Println("Works!")
	} else {
		t.Errorf("Wrong value %v!", result)
	}
}

func TestDetermineHandTypeWithJ(t *testing.T) {
	input := map[string]HandType{
		"2KT54": HIGH_CARDS,
		"9TKJ8": ONE_PAIR,
		"J3JQ2": THREE_OF_KIND,
		"AAJD9": THREE_OF_KIND,
		"AAJKK": FULL_HOUSE,
		"AAJJD": FOUR_OF_KIND,
		"QQQJ9": FOUR_OF_KIND,
		"KTJJJ": FOUR_OF_KIND,
		"J4444": FIVE_OF_KIND,
		"JJJJ4": FIVE_OF_KIND,
		"AAJJJ": FIVE_OF_KIND,
		"JJJJJ": FIVE_OF_KIND,
	}
	for hand, _type := range input {
		result := DetermineHandTypeWithJ(hand)
		if _type == result {
			fmt.Printf("For %s result is correct!\n", hand)
		} else {
			t.Errorf("For %s result is incorrect: %v!\n", hand, result)
		}
	}
}
