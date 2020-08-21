package HouseRobberii

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	nums1 := make([]int, len(nums)-1)
	nums2 := make([]int, len(nums)-1)
	for i, _ := range nums {
		if i > 0 && i <= len(nums)-1 {
			nums1[i-1] = nums[i]
		}
		if i < len(nums)-1 {
			nums2[i] = nums[i]
		}
	}
	return max(robrec(nums1), robrec(nums2))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func robrec(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	first := nums[0]
	second := max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		first, second = second, max(first+nums[i], second)
	}
	return second
}
