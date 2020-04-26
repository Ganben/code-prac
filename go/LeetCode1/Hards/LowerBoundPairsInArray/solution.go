package solution

func reversePairs(nums []int) int {
	res := 0
	for i, n := range nums {
		for j := i; j < len(nums); j++ {
			if nums[j] < n {
				res++
			}
		}
	}
	return res
}
