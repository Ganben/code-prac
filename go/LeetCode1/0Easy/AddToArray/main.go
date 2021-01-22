package main

import "fmt"

func addToArrayForm(A []int, k int) []int {
	res, carry := 0, 0
	arr2 := []int{}
	for k > 0 {
		arr2 = append(arr2, k%10)
		k /= 10
	}
	if len(A) < len(arr2) {
		d := len(arr2) - len(A)
		for i := 0; i < d; i++ {
			A = append([]int{0}, A...)
		}
	}
	for i := 0; i < len(A); i++ {
		if i < len(arr2) {
			cur := A[len(A)-1-i] + arr2[i] + carry
			res = cur % 10
			carry = cur / 10
			A[len(A)-1-i] = res

		} else {
			cur := A[len(A)-1-i] + carry
			res = cur % 10
			carry = cur / 10
			A[len(A)-1-i] = res

		}
		if i == len(A)-1 && carry == 1 {
			A = append([]int{1}, A...)
			break
		}
	}
	return A
}

func main() {
	a := []int{1, 2, 0, 0}
	b := 34
	fmt.Printf("%v\n", addToArrayForm(a, b))
	a = []int{9, 9, 9, 1}
	b = 9
	fmt.Printf("%v\n", addToArrayForm(a, b))
	a = []int{0}
	b = 23
	fmt.Printf("%v\n", addToArrayForm(a, b))
}
