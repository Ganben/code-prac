package main

func getMaximumGenerated(n int) int {
	ans := 0
	if n == 0 {
		return 0
	}
	nums := make([]int, n+1)
	nums[1] = 1
	for i := 2; i <= n; i++ {
		nums[i] = nums[i/2] + i%2*nums[i/2+1]
	}
	for _, v := range nums {
		ans = max(ans, v)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
