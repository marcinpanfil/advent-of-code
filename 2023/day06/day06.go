package aoc2023

func CountNumberOfBestWays(input map[int]int) int {
	total := 1
	for time, distance := range input {
		count := 0
		for i := 0; i <= time; i++ {
			boatDist := i * (time - i)
			if boatDist > distance {
				count++
			}
		}
		total *= count
	}
	return total
}
