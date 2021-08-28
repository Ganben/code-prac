package main

func runningSum(nums []int) []int {
	ans := []int{}
	for i, _ := range nums {
		if i == 0 {
			ans = append(ans, nums[0])
		} else {
			ans = append(ans, nums[i]+ans[i-1])
		}

	}
	return ans
}
