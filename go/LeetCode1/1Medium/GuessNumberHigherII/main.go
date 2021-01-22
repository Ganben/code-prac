package main

import (
	"fmt"
	"math"
)

func getMoneyAmount(n int) int {
	dp := [][]int{}
	for i := 0; i < n+1; i++ {
		row := []int{}
		for j := 0; j < n+1; j++ {
			row = append(row, 0)
		}
		dp = append(dp, row)
	}

	for l := 2; l <= n; l++ {
		for start := 1; start <= n-l+1; start++ {
			minres := math.MaxInt32
			for piv := start; piv < start+l-1; piv++ {
				res := piv + max(dp[start][piv-1], dp[piv+1][start+l-1])
				minres = min(res, minres)
			}
			dp[start][start+l-1] = minres
		}
	}
	return dp[1][n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	fmt.Printf("%v", getMoneyAmount(10))
}
