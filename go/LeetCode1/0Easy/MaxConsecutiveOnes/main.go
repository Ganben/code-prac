package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	pstart, maxLen := -1, 0
	l := 0
	for i := 0; i < len(nums); i++ {

		if nums[i] == 1 && pstart != -1 {
			l++
			// continue
		} else if nums[i] != 1 && pstart != -1 {
			pstart = -1
			if l > maxLen {
				maxLen = l
			}
			l = 0
		} else if nums[i] == 1 && pstart == -1 {
			pstart = i
			l = 1
		}
		if l > maxLen {
			maxLen = l
		}
	}
	//
	return maxLen
}

func main() {
	arr := []int{1, 1, 0, 1, 1, 1}
	fmt.Printf("%v\n", findMaxConsecutiveOnes(arr))
}
