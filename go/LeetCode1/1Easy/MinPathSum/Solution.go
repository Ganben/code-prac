package solution

func minPathSum(grid [][]int) int {
	for i := len(grid) - 1; i >= 0; i-- {
		for j := len(grid[0]) - 1; j >= 0; j-- {
			rN := -1
			dN := -1
			if i+1 <= len(grid)-1 {
				dN = grid[i+1][j]
			}
			if j+1 <= len(grid[0])-1 {
				rN = grid[i][j+1]
			}

			if rN >= 0 && dN >= 0 {
				grid[i][j] = grid[i][j] + min(rN, dN)
			} else if rN >= 0 && dN < 0 {
				grid[i][j] = grid[i][j] + rN
			} else if rN < 0 && dN >= 0 {
				grid[i][j] = grid[i][j] + dN
			} else {
				grid[i][j] = grid[i][j]
			}
		}
	}
	return grid[0][0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
