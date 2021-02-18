package main

import "fmt"

func minKBitFlips(A []int, K int) int {
	ans := 0
	n := len(A)
	diff := make([]int, n+1)
	revCnt := 0
	for i, v := range A {
		revCnt ^= diff[i]
		if v == revCnt {
			if i+K > n {
				return -1
			}
			ans++
			revCnt ^= 1
			diff[i+K] ^= 1
		}
	}
	return ans
}

func main() {
	arr := []int{0, 1, 0}
	k := 1
	fmt.Printf("%v\n", minKBitFlips(arr, k))
}
