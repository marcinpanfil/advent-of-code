package array

import "slices"

func AppendIfNotPresent(data []int, inputs ...int) []int {
	for _, input := range inputs {
		if !slices.Contains(data, input) {
			return append(data, input)
		}
	}
	return data
}
