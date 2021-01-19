package main

import "fmt"

func superPow(a int, b []int) int {
	base := 1337
	if len(b) == 0 {
		return 1
	}
	//
	last := b[len(b)-1]
	b = b[:len(b)-1]
	part1 := myPow(a, last)
	part2 := myPow(superPow(a, b), 10)
	return (part1 * part2) % base
}

func myPow(a, k int) int {
	base := 1337
	a = a % base
	res := 1
	for i := 0; i < k; i++ {
		res *= a
		res %= base
	}
	return res
}

func main() {
	a := 2
	b := []int{1, 0}
	fmt.Printf("%v", superPow(a, b))
}
