package maxsubarray

var dp []int

func maxSubArray(nums []int) int {
	// find max
	// maxvalue = nums[0]
	// middle = 0
	// searchmax(nums)
	dp = []int{0, 0, 0}
	dpAdd(nums)
	return dp[0]

}

func dpAdd(subarray []int) {

	for i, n := range subarray {
		// fmt.Printf("visit %v:%v\n", i, n)
		if i == 0 {
			dp[0] = n
			dp[1] = n
			continue
		}
		// fmt.Printf("state %v: %v , %v\n", i, middle, maxvalue)
		// n < 0
		dp[2] = max(dp[1]+n, n)
		dp[1] = dp[2]
		dp[0] = max(dp[0], dp[2])

	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
