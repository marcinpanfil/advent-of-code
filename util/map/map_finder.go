package maps

import "math"

func FindMaxInIntValues(registers map[string]int) int {
	max := math.MinInt64
	for _, v := range registers {
		if v > max {
			max = v
		}
	}
	return max
}
