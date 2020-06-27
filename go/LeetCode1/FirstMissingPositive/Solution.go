package Solution

import "fmt"

func firstMissingPositive(nums []int) int {
	i := 1
	for _, c := range nums {
		if c > 0 && c == i {
			i += 1
		}
	}
	return i
}

func firstMissingPositive2(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}
	for i := 0; i < n; i++ {
		num := abs(nums[i])
		if num <= n {
			fmt.Println(num - 1)
			nums[num-1] = -abs(nums[num-1])
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
