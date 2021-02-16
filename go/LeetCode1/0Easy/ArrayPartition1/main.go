package main

import (
	"fmt"
	"sort"
)

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	ans := 0
	for i := 0; i < len(nums); i += 2 {
		ans += nums[i]
	}
	return ans
}

func main() {
	arr := []int{1, 4, 3, 2}
	fmt.Printf("%v\n", arrayPairSum(arr))
}
