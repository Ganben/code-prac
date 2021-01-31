package SlidingWindowMaximum

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n*k == 0 {
		return []int{}
	}
	if k == 1 {
		return nums
	}

	left := make([]int, n)
	left[0] = nums[0]
	right := make([]int, n)
	right[n-1] = nums[n-1]
	for i := 1; i < n; i++ {
		if i%k == 0 {
			left[i] = nums[i]
		} else {
			left[i] = max(left[i-1], nums[i])
		}
		j := n - i - 1
		if (j+1)%k == 0 {
			right[j] = nums[j]
		} else {
			right[j] = max(right[j+1], nums[j])
		}
	}
	output := make([]int, n-k+1)
	for i := 0; i < n-k+1; i++ {
		output[i] = max(left[i+k-1], right[i])
	}
	return output
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
