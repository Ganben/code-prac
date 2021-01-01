package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	hMap := map[int]int{}
	for _, v := range nums1 {
		_, exist := hMap[v]
		if !exist {
			hMap[v] = 1
		}
	}
	for _, v := range nums2 {
		val, exist := hMap[v]
		if !exist {
			hMap[v] = 3
		} else if val == 1 {
			hMap[v] = 2
		}
	}
	ans := []int{}
	for k, v := range hMap {
		if v == 2 {
			ans = append(ans, k)
		}
	}
	return ans
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2, 3, 3}
	fmt.Printf("%v", intersection(nums1, nums2))
}
