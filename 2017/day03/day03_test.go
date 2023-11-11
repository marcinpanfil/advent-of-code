package aoc2017

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	fmt.Println(ConstructSquareSequentially(325489)) // 552
	fmt.Println(ConstructSquareAdjacent(325489))     // 330785
}

func TestConstructSquareSequentially(t *testing.T) {
	data := map[int]int{
		1:    0,
		12:   3,
		23:   2,
		32:   5,
		42:   5,
		47:   4,
		1024: 31,
	}

	for k, v := range data {
		result := ConstructSquareSequentially(k)
		if v == result {
			fmt.Printf("For value %v there is correct result %v\n", k, v)
		} else {
			t.Errorf("For value %v there is incorrect result %v\n", k, result)
		}
	}
}

func TestConstructSquareAdjacent(t *testing.T) {
	data := map[int]int{
		750: 806,
		50:  54,
	}

	for k, v := range data {
		result := ConstructSquareAdjacent(k)
		if v == result {
			fmt.Printf("For value %v there is correct result %v\n", k, v)
		} else {
			t.Errorf("For value %v there is incorrect result %v\n", k, result)
		}
	}
}
