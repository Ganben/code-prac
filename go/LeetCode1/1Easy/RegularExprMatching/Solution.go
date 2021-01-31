package Solution

func isMatch(s string, p string) bool {
	//
	dp := make([][]bool, len(s)+1)
	for i := 0; i < len(s) + 1; i++ {
		dp[i] = make([]bool, len(p) + 1)
	}
	dp[len(s)][len(p)] = true

	for i:= len(s); i >= 0; i-- {
		for j:= len(p) - 1; j >= 0; j-- {
			//
			fm := false
			if i < len(s) && (s[i] == p[j] || p[j] == '.') {
				fm = true
			}
			//
			if (j + 1) < len(p) && p[j+1] == '*' {
				//
				dp[i][j] = dp[i][j+2] || (fm && dp[i+1][j])
			} else {
				dp[i][j] = fm && dp[i+1][j+1]
			}
		}
	}
	return dp[0][0]
}