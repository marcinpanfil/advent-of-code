package aoc2017

func SpinlockAlgorithm(steps int) int {
	state := []int{0}
	curPos := 0
	for i := 1; i <= 2017; i++ {
		curPos = (curPos+steps)%len(state) + 1
		if curPos == 0 {
			state = append([]int{i}, state...)
		} else if curPos == len(state) {
			state = append(state, i)
		} else {
			state = append(state[:curPos], append([]int{i}, state[curPos:]...)...)
		}
	}
	return state[curPos+1]
}

func SpinlockAlgorithm50kk(steps int) int {
	result := 0
	curPos := 1
	curLen := 2
	for i := 2; i <= 50_000_000; i++ {
		curPos = (curPos+steps)%curLen + 1
		// because 0 is always at pos 0, so it's only important to know what is at pos 1
		if curPos == 1 {
			result = i
		}
		curLen++
	}
	return result
}
