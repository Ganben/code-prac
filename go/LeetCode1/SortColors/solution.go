package SortColors

func sortColors(nums []int) {
	curr, p0, p1 := 0, 0, len(nums)-1
	for curr <= p1 {
		switch nums[curr] {
		case 0:
			nums[p0], nums[curr] = nums[curr], nums[p0]
			curr++
			p0++
		case 1:
			curr++
		case 2:
			nums[curr], nums[p1] = nums[p1], nums[curr]
			p1--
		}
	}

}
