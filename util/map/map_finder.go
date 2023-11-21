package maps

import "math"

func FindMaxInIntValues(data map[string]int) int {
	max := math.MinInt64
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	return max
}
