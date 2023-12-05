package aoc2023

import (
	"advent-of-code/util/file"
	"fmt"
	"strings"
	"testing"
)

var INPUT = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

func TestSolution(t *testing.T) {
	input := file.ReadInput()
	fmt.Println(FindMinLocation(input)) // 261668924
	// fmt.Println(FindMinLocationForRanges(input)) // 24261545
}

func TestFindMinLocationPart1(t *testing.T) {
	result := FindMinLocation(strings.Split(INPUT, "\n"))
	if result == 35 {
		fmt.Println("Correct!")
	} else {
		t.Errorf("Wrong value %v!", result)
	}
}

func TestFindMinLocationPart2(t *testing.T) {
	result := FindMinLocationForRanges(strings.Split(INPUT, "\n"))
	if result == 46 {
		fmt.Println("Correct!")
	} else {
		t.Errorf("Wrong value %v!", result)
	}
}
