package main

import "fmt"

func fairCandySwap(A []int, B []int) []int {
	// p1, p2 := 0, 0
	ans := []int{}
	sumA, sumB := 0, 0
	for i := 0; i < len(A); i++ {
		sumA += A[i]
	}
	for i := 0; i < len(B); i++ {
		sumB += B[i]
	}
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			if arsum(sumA, A[i], B[j]) == arsum(sumB, B[j], A[i]) {
				ans = append(ans, A[i])
				ans = append(ans, B[j])
				return ans
			}
		}
	}
	return ans
}

func arsum(suM int, r int, val int) int {
	return suM - r + val
}

func main() {
	a := []int{1, 1}
	b := []int{2, 2}
	fmt.Printf("%v\n", fairCandySwap(a, b))
}
