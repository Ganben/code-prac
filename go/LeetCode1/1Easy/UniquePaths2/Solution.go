package solution

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	for i := len(obstacleGrid) - 1; i >= 0; i-- {
		for j := len(obstacleGrid[0]) - 1; j >= 0; j-- {
			if i == len(obstacleGrid)-1 && j == len(obstacleGrid[0])-1 {
				if obstacleGrid[i][j] == 0 {
					obstacleGrid[i][j] = -1
					continue
				}
			}
			rN := 0
			dN := 0
			if i+1 <= len(obstacleGrid)-1 {
				if obstacleGrid[i+1][j] != 1 {
					dN = obstacleGrid[i+1][j]
				}
			}
			if j+1 <= len(obstacleGrid[0])-1 {
				if obstacleGrid[i][j+1] != 1 {
					rN = obstacleGrid[i][j+1]
				}
			}
			if obstacleGrid[i][j] != 1 {
				obstacleGrid[i][j] = rN + dN
			}

		}
	}
	if obstacleGrid[0][0] <= 0 {
		return -1 * obstacleGrid[0][0]
	}
	return 0
}
