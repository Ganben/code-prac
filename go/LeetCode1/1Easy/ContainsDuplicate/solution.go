package ContainsDuplicate

import "sort"

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i, _ := range nums {
		if i < len(nums)-1 {
			if nums[i] == nums[i+1] {
				return true
			}
		}
	}
	return false
}
