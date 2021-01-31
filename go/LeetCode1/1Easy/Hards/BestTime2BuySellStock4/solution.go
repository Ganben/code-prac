package BestTime2BuySellStock4

import "math"

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if k > n/2 {
		return maxProfitInf(prices)
	}
	dp := [][][]int{}
	for i := 0; i < n; i++ {
		subMatrix := [][]int{}
		for i := 0; i < k+1; i++ {
			row := []int{0, math.MinInt32}
			subMatrix = append(subMatrix, row)

		}
		dp = append(dp, subMatrix)
	}

	for i := 0; i < n; i++ {
		for kk := k; kk >= 1; kk-- {
			if i == 0 {
				dp[i][kk][0] = prices[i]

			} else {

				dp[i][kk][0] = max(dp[i-1][kk][0], dp[i-1][kk][1]+prices[i])
				dp[i][kk][1] = max(dp[i-1][kk][1], dp[i-1][kk-1][0]-prices[i])
			}
		}
	}
	return dp[n-1][k][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxProfitInf(prices []int) int {
	n := len(prices)
	dpI0 := 0
	dpI1 := math.MinInt32
	dpIp := 0
	for i := 0; i < n; i++ {
		tmp := dpI0
		dpI0 = max(dpI0, dpI1+prices[i])
		dpI1 = max(dpI1, dpIp-prices[i])
		dpIp = tmp
	}
	return dpI0
}
