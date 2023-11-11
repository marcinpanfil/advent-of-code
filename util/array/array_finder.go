package array

import "math"

func MaxIndexAndValue(input []int) (int, int) {
	if len(input) < 1 {
		panic("Incorrect lenght of the array")
	}
	maxValue := math.MinInt32
	index := -1
	for i, v := range input {
		if v > maxValue {
			maxValue = v
			index = i
		}
	}
	return index, maxValue
}
