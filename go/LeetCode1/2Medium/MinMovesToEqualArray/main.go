package minmovestoequalarray

import "sort"

func minMoves2(nums []int) int {
	length := len(nums)
	sort.Ints(nums)
	num := nums[length/2]
	res := 0
	for i := 0; i < length; i++ {
		res += getAbs(num, nums[i])
	}
	return res
}

func getAbs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
