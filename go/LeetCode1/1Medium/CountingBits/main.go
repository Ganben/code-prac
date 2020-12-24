package main

import "fmt"

func countBits(num int) []int {
	ans := []int{}
	for i := 0; i <= num; i++ {
		ans = append(ans, testBits(i))
	}
	//
	return ans
}

func testBits(num int) int {
	//
	if num == 0 {
		return 0
	}
	if num%2 == 1 {
		return 1 + testBits(num>>1)
	} else {
		return testBits(num >> 1)
	}
}

func main() {
	fmt.Printf("%v", countBits(5))
}
