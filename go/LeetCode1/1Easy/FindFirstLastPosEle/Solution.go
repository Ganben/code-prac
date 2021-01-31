package Solution

func searchRange(nums []int, target int) []int {
	targetRange := []int{-1, -1}

	for i := 0; i < len(nums); i ++ {
		if nums[i] == target {
			targetRange[0] = i
			break
		}
	}
	if targetRange[0] == -1 {
		return targetRange
	}

	for j := len(nums) - 1; j >= 0; j -- {
		if nums[j] == target {
			targetRange[1] = j
			break
		}
	}
	return targetRange
}