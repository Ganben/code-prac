package main

import "fmt"

func getSum(a, b int) int {
	for b != 0 {
		sum := a ^ b
		carry := (a & b) << 1
		a = sum
		b = carry
	}
	return a
}

func main() {
	fmt.Printf("%v", getSum(1, 2))
}
