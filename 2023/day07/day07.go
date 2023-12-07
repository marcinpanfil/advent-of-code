package aoc2023

import (
	"advent-of-code/util/array"
	"sort"
	"strconv"
	"strings"
)

type HandType int8

const (
	FIVE_OF_KIND  HandType = 6
	FOUR_OF_KIND  HandType = 5
	FULL_HOUSE    HandType = 4
	THREE_OF_KIND HandType = 3
	TWO_PAIR      HandType = 2
	ONE_PAIR      HandType = 1
	HIGH_CARDS    HandType = 0
)

type Hand struct {
	hand      string
	_type     HandType
	bid       int
	typeWithJ HandType
}

var CARDS_ORDER = []rune{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}
var CARDS_ORDER_WITH_J = []rune{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}

type ByType []Hand
type ByTypeWithJ []Hand

func (hands ByType) Len() int { return len(hands) }
func (hands ByType) Less(i, j int) bool {
	first := hands[i]
	second := hands[j]
	if first._type < second._type {
		return true
	} else if first._type == second._type {
		firstHand := first.hand
		secondHand := second.hand
		for k, card := range firstHand {
			card1 := array.IndexOf(CARDS_ORDER, card)
			card2 := array.IndexOf(CARDS_ORDER, rune(secondHand[k]))
			if card1 < card2 {
				return true
			} else if card1 > card2 {
				return false
			}
		}
		panic("The same hand!")
	} else {
		return false
	}
}
func (hands ByType) Swap(i, j int) { hands[i], hands[j] = hands[j], hands[i] }

func (hands ByTypeWithJ) Len() int { return len(hands) }
func (hands ByTypeWithJ) Less(i, j int) bool {
	first := hands[i]
	second := hands[j]
	if first.typeWithJ < second.typeWithJ {
		return true
	} else if first.typeWithJ == second.typeWithJ {
		firstHand := first.hand
		secondHand := second.hand
		for k, card := range firstHand {
			card1 := array.IndexOf[rune](CARDS_ORDER_WITH_J, card)
			card2 := array.IndexOf[rune](CARDS_ORDER_WITH_J, rune(secondHand[k]))
			if card1 < card2 {
				return true
			} else if card1 > card2 {
				return false
			}
		}
		panic("The same hand!")
	} else {
		return false
	}
}
func (hands ByTypeWithJ) Swap(i, j int) { hands[i], hands[j] = hands[j], hands[i] }

func CalculateTotalWinnings(input []string, joker bool) int {
	hands := ParseInput(input)
	if joker {
		sort.Sort(ByTypeWithJ(hands))
	} else {
		sort.Sort(ByType(hands))
	}
	result := 0
	for i, hand := range hands {
		result += (i + 1) * hand.bid
	}
	return result
}

func ParseInput(input []string) []Hand {
	result := make([]Hand, len(input))
	for i, line := range input {
		data := strings.Split(line, " ")
		handStr := data[0]
		bid, _ := strconv.Atoi(data[1])
		handType := DetermineHandType(handStr)
		handTypeWithJ := DetermineHandTypeWithJ(handStr)
		result[i] = Hand{hand: handStr, bid: bid, _type: handType, typeWithJ: handTypeWithJ}
	}
	return result
}

func DetermineHandTypeWithJ(hand string) HandType {
	jCount := strings.Count(hand, "J")
	if jCount > 0 {
		cards := map[rune]int{}
		for _, card := range hand {
			cards[card] = cards[card] + 1
		}
		if len(cards) == 1 || len(cards) == 2 {
			return FIVE_OF_KIND
		} else if len(cards) == 3 {
			for _, count := range cards {
				if count == 3 || (count == 2 && jCount == 2) {
					return FOUR_OF_KIND
				} else if count == 2 && jCount == 1 {
					return FULL_HOUSE
				}
			}
		} else if len(cards) == 4 {
			return THREE_OF_KIND
		} else {
			return ONE_PAIR
		}
		panic("Wrong hand! No option fit!")
	} else {
		return DetermineHandType(hand)
	}
}

func DetermineHandType(hand string) HandType {
	cards := map[rune]int{}
	for _, card := range hand {
		cards[card] = cards[card] + 1
	}
	if len(cards) == 1 {
		return FIVE_OF_KIND
	} else if len(cards) == 2 {
		for _, count := range cards {
			if count == 4 || count == 1 {
				return FOUR_OF_KIND
			} else if count == 2 || count == 3 {
				return FULL_HOUSE
			} else {
				panic("Wrong hand!")
			}
		}
	} else if len(cards) == 3 {
		for _, count := range cards {
			if count == 3 {
				return THREE_OF_KIND
			} else if count == 2 {
				return TWO_PAIR
			}
		}
	} else if len(cards) == 4 {
		return ONE_PAIR
	} else {
		return HIGH_CARDS
	}
	panic("Wrong hand! No option fit!")
}
