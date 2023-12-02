package aoc2017

var GEN_A = int64(16807)
var GEN_B = int64(48271)
var DIV = int64(2147483647)

func CalculateValueOfGen(genA int64, genB int64) int {
	counter := 0
	for i := 0; i < 40_000_000; i++ {
		genA *= GEN_A
		genB *= GEN_B
		genA %= DIV
		genB %= DIV
		if int16(genA) == int16(genB) {
			counter++
		}
	}
	return counter
}

func CalculateValueOfPickyGen(genA int64, genB int64) int {
	counter := 0
	for i := 0; i < 5_000_000; i++ {
		genA *= GEN_A
		genB *= GEN_B
		genA %= DIV
		genB %= DIV

		for genA%4 != 0 {
			genA *= GEN_A
			genA %= DIV
		}

		for genB%8 != 0 {
			genB *= GEN_B
			genB %= DIV
		}

		if int16(genA) == int16(genB) {
			counter++
		}
	}
	return counter
}
