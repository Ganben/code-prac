package Solution

func isValidSudoku(board [][]byte) bool {
	n := len(board)
	rectSets := make([]int16, n*n/9)
	for i := 0; i < n; i ++ {
		rowSet, colSet := int16(0), int16(0)
		for j := 0; j < n; j ++ {
			if num := board[i][j]; num != '.' {
				numBit := int16(1 << (num - '0'))
				if rowSet&numBit > 0 {
					return false
				}
				rowSet |= numBit
				if rectSets[(n/3)*(i/3) + j/3]&numBit > 0 {
					return false
				}
				rectSets[(n/3)*(i/3) + j/3] |= numBit
			}
			if num := board[j][i]; num != '.' {
				numBit := int16(1 << (num - '0'))
				if colSet&numBit > 0 {
					return false
				}
				colSet ^= numBit
			}
		}
	}
	return true
}