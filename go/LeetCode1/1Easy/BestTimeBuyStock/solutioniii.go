package BestTimeBuyStock

import "math"

func maxProfit3(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := make([]int, 5)
	dp[0] = 0
	dp[1] = -prices[0]
	dp[3] = math.MinInt32
	for _, v := range prices {
		dp[0] = 0
		dp[1] = max(dp[1], -1*v)
		dp[2] = max(dp[2], dp[1]+v)
		dp[3] = max(dp[3], dp[2]-v)
		dp[4] = max(dp[4], dp[3]+v)
	}
	return max3(dp[0], dp[2], dp[4])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func max3(x, y, z int) int {
	if x > y {
		return max(x, z)
	} else {
		return max(y, z)
	}
}
