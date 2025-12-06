package utils

func Transpose[T any](matrix [][]T) [][]T {
	if len(matrix) == 0 {
		return matrix
	}

	rows := len(matrix)
	cols := len(matrix[0])

	// Create new matrix with swapped dimensions
	result := make([][]T, cols)
	for i := range result {
		result[i] = make([]T, rows)
	}

	// Swap rows and columns
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[j][i] = matrix[i][j]
		}
	}

	return result
}
