package array

import (
	"slices"
)

func AppendIfNotPresent[T comparable](data []T, inputs ...T) []T {
	for _, input := range inputs {
		if !slices.Contains(data, input) {
			data = append(data, input)
		}
	}
	return data
}
