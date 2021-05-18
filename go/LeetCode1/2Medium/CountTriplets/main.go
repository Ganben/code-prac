package main

import "fmt"

func countTriplets(arr []int) int {
	ans := 0
	n := len(arr)
	s := make([]int, n+1)
	for i, val := range arr {
		s[i+1] = s[i] ^ val
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j; k < n; k++ {
				if s[i] == s[k+1] {
					ans++
				}
			}
		}
	}
	return ans
}

func main() {
	arr := []int{2, 3, 1, 6, 7}
	fmt.Printf("%v\n", countTriplets(arr))
}
