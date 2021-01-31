package Solution

func jump(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	dp := make([]int, n)
	for i := 0; i < len(nums); i++ {
		dp[i] = n + 99
	}
	dp[0] = 0
	for i := 0; i < n-1; i++ {
		for j := 1; j <= nums[i]; j++ {
			if i+j < n {
				dp[i+j] = Min(dp[i+j], dp[i]+1)
			}
		}
	}
	return dp[n-1]
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
