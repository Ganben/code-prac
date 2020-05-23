package RemoveDuplFromSortedArray

func removeDuplicates(nums []int) int {
	i, count := 1, 1
	for i < len(nums) {
		if nums[i] == nums[i-1] {
			count++
			if count > 2 {
				tmp1 := nums[:i]
				tmp2 := nums[i+1:]
				nums = []int{}
				nums = append(nums, tmp1...)
				nums = append(nums, tmp2...)

			}

		} else {
			count = 1
		}
		i++
	}
	return len(nums)
}
