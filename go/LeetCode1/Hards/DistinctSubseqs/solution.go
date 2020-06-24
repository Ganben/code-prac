package DistinctSubseqs

func numDistinct(s string, t string) int {
	n1 := len(s)
	n2 := len(t)
	dp := [][]int{}
	for i := 0; i < n2+1; i++ {
		row := make([]int, n1+1)
		dp = append(dp, row)

	}
	for j := 0; j < n1+1; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < n2+1; i++ {
		for j := 1; j < n1+1; j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[n2][n1]
}
