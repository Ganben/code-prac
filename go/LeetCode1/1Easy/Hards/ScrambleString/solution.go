package solution

func isScramble1(s1, s2 string) bool {
	n1 := len(s1)
	n2 := len(s2)
	if n1 != n2 {
		return false
	}
	var dp [][][]bool
	dp = [][][]bool{}
	row2 := [][]bool{}
	for i := 0; i < n1; i++ {
		row2 = [][]bool{}
		for j := 0; j < n1; j++ {
			row := make([]bool, n1+1)
			row[1] = s1[i] == s2[j]
			row2 = append(row2, row)
		}
		dp = append(dp, row2)
	}
	for l := 2; l < n1; l++ {
		for i := 0; i < n1-l; i++ {
			for j := 0; j < n1-l; j++ {
				for k := 1; k < l-1; k++ {
					//case 1
					if dp[i][j][k] && dp[i+k][j+k][l-k] {
						dp[i][j][l] = true
						break
					}
					// case 2
					if dp[i][j+l-k][k] && dp[i+k][j][l-k] {
						dp[i][j][l] = true
						break
					}
				}
			}
		}
	}
	return dp[0][0][n1]
}

func isScramble(s1, s2 string) bool {

}
