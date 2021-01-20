package main

import (
	"fmt"
	"math"
)

func maximumProduct(nums []int) int {
	r := max3(nums)

	ans := 1
	for _, v := range r {
		ans *= v
	}
	return ans
}

func max3(nums []int) []int {
	ans := []int{}
	if len(nums) < 3 {
		return ans
	}
	//
	// max one
	maxNum, maxId := 1-math.MaxInt32, 0
	//
	for i, v := range nums {
		if v > maxNum {
			maxNum = v
			maxId = i
		}
	}
	ans = append(ans, maxNum)
	b1 := nums[:maxId]
	b1 = append(b1, nums[maxId+1:]...)
	if maxNum > 0 {
		ans = append(ans, max2(b1))
	} else {
		ans = append(ans, min2(b1))
	}

	return ans
}

func min2(nums []int) int {
	minP := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]*nums[j] < minP {
				minP = nums[i] * nums[j]
			}
		}
	}
	return minP
}

func max2(nums []int) int {
	maxP := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]*nums[j] > maxP {
				maxP = nums[i] * nums[j]
			}
		}
	}
	return maxP
}

func main() {
	b := []int{1, 2, 3, 4, 5}
	fmt.Printf("%v", maximumProduct(b))
}
