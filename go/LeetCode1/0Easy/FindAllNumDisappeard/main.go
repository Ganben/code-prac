package main

import "fmt"

func findDisappearedNumber(nums []int) []int {
	n := len(nums)
	ans := []int{}
	for _, v := range nums {
		v = (v - 1) % n
		nums[v] += n
	}
	for i, v := range nums {
		if v <= n {
			ans = append(ans, i+1)
		}
	}
	return ans
}

func main() {
	arr := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Printf("%v\n", findDisappearedNumber(arr))
}
