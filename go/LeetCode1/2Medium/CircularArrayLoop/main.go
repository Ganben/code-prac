package main

import "fmt"

func circularArrayLoop(nums []int) bool {
	IFMAX := 1000
	n := len(nums)
	for index, num := range nums {
		if num > -IFMAX && num < IFMAX {
			lastIndex := index
			i := index
			for nums[i] > 0 && nums[i] <= IFMAX {
				lastIndex = i
				//
				i = (nums[i] + i) % n
				nums[lastIndex] = IFMAX + index + 1
			}
			//
			if lastIndex != i && nums[i] == IFMAX+index+1 {
				return true
			}
			//
			for nums[i] >= -IFMAX && nums[i] < 0 {
				lastIndex = i
				//
				i = (n - (-nums[i] % n) + i) % n
				nums[lastIndex] = -IFMAX - index - 1
			}
			//
			if lastIndex != i && nums[i] == -IFMAX-index-1 {
				return true
			}

		}
	}
	return false
}

func main() {
	ns := []int{2, -1, 1, 2, 2}
	fmt.Printf("%v\n", circularArrayLoop(ns))
}
