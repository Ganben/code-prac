package UglyNumber2

func nthUglyNumber(n int) int {
	nums := []int{1}
	i2, i3, i5 := 0, 0, 0
	for i := 1; i <= 1690; i++ {
		ugly := min(nums[i2]*2, nums[i3]*3, nums[i5]*5)
		nums = append(nums, ugly)
		if ugly == nums[i2]*2 {
			i2++
		}
		if ugly == nums[i3]*3 {
			i3++
		}
		if ugly == nums[i5]*5 {
			i5++
		}
	}
	// fmt.Printf("%v", nums)
	return nums[n-1]
}

func min(x, y, z int) int {
	if x < y {
		if x < z {
			return x
		}
		return z
	}
	if y > z {
		return z
	}
	return y
}
