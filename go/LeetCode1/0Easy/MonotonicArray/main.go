package main

import "fmt"

func isMonotonic(A []int) bool {
	if len(A) <= 2 {
		return true
	}
	sign := A[1] - A[0]
	for i := 1; i < len(A); i++ {
		s := A[i] - A[i-1]
		if sign == 0 && s != 0 {
			sign = s
		}
		if s*sign < 0 {
			return false
		}
	}
	return true
}

func main() {
	i1 := []int{2, 2, 2, 1, 4, 5}
	fmt.Printf("%v\n", isMonotonic(i1))
}
