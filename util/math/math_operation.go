package math

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(integers ...int) int {
	result := integers[0] * integers[1] / GCD(integers[0], integers[1])

	for i := 2; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func GaussianElimination(matrix [][]float64) [][]float64 {
	rows := len(matrix)
	cols := len(matrix[0])

	for i := 0; i < rows; i++ {
		pivot := matrix[i][i]
		for j := i; j < cols; j++ {
			matrix[i][j] /= pivot
		}

		for k := 0; k < rows; k++ {
			if k != i {
				scale := matrix[k][i]
				for j := i; j < cols; j++ {
					matrix[k][j] -= scale * matrix[i][j]
				}
			}
		}
	}
	return matrix
}
