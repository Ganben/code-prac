package main

import (
	"fmt"
)

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	if len(nums) <= k {
		return float64(sumArr(nums)) / float64(k)
	}
	sum = sumArr(nums[:k])
	maxSum := sum
	for i := 1; i <= len(nums)-k; i++ {
		sum -= nums[i-1]
		sum += nums[i+k-1]
		if maxSum < sum {
			maxSum = sum
		}
		fmt.Printf("%v\n", sum)
	}
	return float64(maxSum) / float64(k)
}

func sumArr(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		ans += nums[i]
	}
	return ans
}

func main() {
	arr := []int{0, 1, 1, 3, 3}
	fmt.Printf("%v\n", findMaxAverage(arr, 4))
}
