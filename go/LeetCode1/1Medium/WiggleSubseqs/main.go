package main

import "fmt"

func wiggleMaxLength(nums []int) int {
	down, up := 1, 1
	if len(nums) == 0 {
		return 0
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			up = down + 1
		} else if nums[i] < nums[i-1] {
			down = up + 1
		}
	}
	if down > up {
		return down
	}
	return up
}

func main() {
	arr := []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}
	fmt.Printf("%v", wiggleMaxLength(arr))
}
