package array

func GenerateSequence(start int, end int) []int {
	ints := make([]int, end-start+1)
	for i := range ints {
		ints[i] = start + i
	}
	return ints
}
