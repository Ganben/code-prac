package EditDistance

func minDistance(word1 string, word2 string) int {
	n := len(word1)
	m := len(word2)

	if n*m == 0 {
		return n + m
	}

	dp := [][]int{}
	for i := 0; i <= n; i++ {
		row := make([]int, m+1)
		row[0] = i
		dp = append(dp, row)
	}
	for i := 0; i <= m; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			left := dp[i-1][j] + 1
			down := dp[i][j-1] + 1
			left_down := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				left_down += 1
			}
			dp[i][j] = min(left, min(down, left_down))
		}
	}
	return dp[n][m]

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
