package main

import (
	"fmt"
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	ans := []int{}
	if len(nums) == 0 {
		return ans
	}
	sort.Ints(nums)
	dp := make([]int, len(nums))
	//

	for i, _ := range nums {
		maxSize := 0
		for k := 0; k <= i; k++ {
			if nums[i]%nums[k] == 0 {
				maxSize = max(maxSize, dp[k])
			}
		}
		maxSize++
		dp[i] = maxSize
	}
	// maxIdx := 0
	maxSize, maxIdx := maxArr(dp)
	currSize, currTail := maxSize, nums[maxIdx]
	for i := maxIdx; i >= 0; i-- {
		if currSize == dp[i] && currTail%nums[i] == 0 {
			ans = append(ans, nums[i])
			currSize--
			currTail = nums[i]
		}
	}
	return reverse(ans)

}

func reverse(arr []int) []int {
	ans := []int{}
	for i := len(arr) - 1; i >= 0; i-- {
		ans = append(ans, arr[i])
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxArr(arr []int) (int, int) {
	idx, m := 0, 0
	for i, v := range arr {
		if v > m {
			idx = i
			v = m
		}
	}
	return m, idx
}

func main() {
	a := []int{1, 2, 3}
	fmt.Printf("%v", largestDivisibleSubset(a))
}
