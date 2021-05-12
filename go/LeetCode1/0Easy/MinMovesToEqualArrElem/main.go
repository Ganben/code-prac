package main

import (
	"fmt"
	"sort"
)


func minMoves(nums []int) int {
	// deduce 1 each el
	sort.Ints(nums)
	ans := 0
	for i:=1; i<len(nums); i++ {
		ans += nums[i] - nums[0]
	}
	return ans
}

func main() {
	i := []int{1, 2, 3}
	fmt.Printf("%v\n", minMoves(i))
}
