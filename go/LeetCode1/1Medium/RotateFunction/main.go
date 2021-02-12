package main

import (
	"fmt"
)

func maxRotateFunction(A []int) int {
	maxRes := arrMultiply(A, 0)
	for i := 1; i < len(A); i++ {
		r := arrMultiply(A, i)
		// fmt.Printf("%v\n", r)
		if maxRes < r {
			maxRes = r
		}
	}
	return maxRes
}

func arrMultiply(A []int, offset int) int {
	sum := 0
	for i, v := range A {
		m := i + offset
		if m >= len(A) {
			m -= len(A)
		}
		// fmt.Printf("+%v*%v>>", m, v)
		sum += m * v
	}
	return sum
}

func main() {
	A := []int{4, 3, 2, 6}
	fmt.Printf("%v\n", maxRotateFunction(A))
}
