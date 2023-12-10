package array

import "slices"

func PositionsEquals(first, second [][2]int) bool {
	for _, pos := range first {
		if !slices.Contains(second, pos) {
			return false
		}
	}
	return true
}
