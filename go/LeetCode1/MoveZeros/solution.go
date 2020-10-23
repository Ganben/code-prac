package MoveZeros

func moveZeroes(nums []int) {
	lastNoneZeroFound := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			lastNoneZeroFound++
			nums[lastNoneZeroFound] = nums[i]
		}
	}
	for i := lastNoneZeroFound; i < len(nums); i++ {
		nums[i] = 0
	}
}
