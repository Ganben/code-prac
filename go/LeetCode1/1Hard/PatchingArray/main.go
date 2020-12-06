package main

import "fmt"

func minPatches(nums []int, n int) int {
	patches, i := 0, 0
	// var miss int64
	miss := 1
	for miss <= n {
		if i < len(nums) {
			if nums[i] <= miss {

				miss += nums[i]
				i++
				continue
			}
		}
		miss += miss
		patches++

	}
	return patches
}

func main() {
	nums := []int{1, 2, 3, 8}
	n := 80
	res := minPatches(nums, n)
	fmt.Printf("%v", res)
}
