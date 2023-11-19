package aoc2017

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	fmt.Println(KnotTyingHash([]int{94, 84, 0, 79, 2, 27, 81, 1, 123, 93, 218, 23, 103, 255, 254, 243}, 256)) // 23715
	fmt.Println(CalculateKnotHash("94,84,0,79,2,27,81,1,123,93,218,23,103,255,254,243"))                      // 541dc3180fd4b72881e39cf925a50253
}

func TestCalculateKnotHash(t *testing.T) {
	data := map[string]string{
		"":         "a2582a3a0e66e6e86e3812dcb672a272",
		"AoC 2017": "33efeb34ea91902bb2f59c9920caa6cd",
		"1,2,3":    "3efbe78a8d82f29979031a4aa0b16a9d",
		"1,2,4":    "63960835bcdc130f0b66d7ff4f6a5a8e",
	}
	for k, v := range data {
		result := CalculateKnotHash(k)
		if result == v {
			fmt.Printf("For %s the value is correct!\n", k)
		} else {
			t.Errorf("For %s the value %s is incorrect!\n", k, result)
		}
	}
}

func TestKnotTyingHash(t *testing.T) {
	result := KnotTyingHash([]int{3, 4, 1, 5}, 5)
	if result == 12 {
		fmt.Println("Correct!")
	} else {
		t.Errorf("Incorrect result: %v", result)
	}
}
