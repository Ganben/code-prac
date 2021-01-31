package MaxGap

import "sort"

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	numsA := sort.IntSlice(nums)
	sort.Sort(numsA)
	maxGap := 0
	for i, v := range numsA {
		if i > 0 {
			maxGap = max(v-nums[i-1], maxGap)
		}

	}
	return maxGap
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
