package main

import (
	"fmt"
	"sort"
)

func longestOnes(A []int, K int) int {
	n := len(A)
	P := make([]int, n+1)
	ans := 0
	for i, v := range A {
		P[i+1] = P[i] + 1 - v
	}
	for right, v := range P {
		left := sort.SearchInts(P, v-K)
		ans = max(ans, right-left)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	arr := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	K := 2
	fmt.Printf("%v\n", longestOnes(arr, K))
}
