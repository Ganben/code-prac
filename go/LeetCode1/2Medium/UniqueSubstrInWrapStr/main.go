package main

func findSubstringInWraproundString(p string) int {
	return dp(p)
}

func dp(p string) int {
	if len(p) < 2 {
		return len(p)
	}
	mapAbc := make([]int, 123)
	dp := make([]int, len(p))
	dp[0] = 1
	mapAbc[p[0]] = 1

	for i := 1; i < len(p); i++ {
		dp[i] = 1
		if p[i]-p[i-1] == 1 || int(p[i])-int(p[i-1]) == -25 {
			dp[i] += dp[i-1]
		}
		mapAbc[p[i]] = getMax(dp[i], mapAbc[p[i]])

	}
	cnt := 0
	for _, n := range mapAbc {
		cnt += n
	}
	return cnt

}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
