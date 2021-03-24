package main

import (
	"fmt"
)

func canPartition(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}
	sum, max := 0, 0
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}
	if sum%2 != 0 {
		return false
	}
	//
	target := sum / 2
	if max > target {
		return false
	}
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}
	//
	for i := 0; i < n; i++ {
		dp[i][0] = true
	}
	dp[0][nums[0]] = true
	for i := 1; i < n; i++ {
		v := nums[i]
		for j := 1; j <= target; j++ {
			if j >= v {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n-1][target]
}

// func arsum(nums []int) int {
// 	ans := 0
// 	for i := 0; i < len(nums); i++ {
// 		ans += nums[i]
// 	}
// 	return ans
// }

func main() {
	ar := []int{1, 5, 11, 5}
	fmt.Printf("%v\n", canPartition(ar))
}
