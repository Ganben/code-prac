package SetMatrixZeros

func setZeroes(matrix [][]int) {
	// iter for zero matrix
	jobStacks := [][]int{}
	if len(matrix) == 0 {
		return
	}
	if len(matrix[0]) == 0 {
		return
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				jobStacks = append(jobStacks, []int{i, j})
			}
		}
	}

	for _, row := range jobStacks {
		// make row 0
		for i := 0; i < len(matrix); i++ {
			matrix[i][row[1]] = 0
		}
		// make col 0
		for j := 0; j < len(matrix[0]); j++ {
			matrix[row[0]][j] = 0
		}
	}

}
