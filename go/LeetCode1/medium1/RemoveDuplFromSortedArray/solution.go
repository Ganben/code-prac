package RemoveDuplFromSortedArray

func removeDuplicates(nums []int) int {
	i, count := 1, 1
	for ; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			count++
			if count > 2 {
				tmp1 := nums[:i]
				tmp2 := nums[i+1:]
				tmp := []int{}
				tmp = append(nums, tmp1...)
				tmp = append(nums, tmp2...)
				nums = tmp
			}

		} else {
			count = 1
		}
	}
	return len(nums)
}
