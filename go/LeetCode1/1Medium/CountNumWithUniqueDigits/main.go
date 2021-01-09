package main

import "fmt"

func countNumbersWithUniqueDigits(n int) int {
	// digits = 0-9
	// digits arr length = n
	// 0 is exception, and 0-9 is all available
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 10
	}

	c := 1
	for i := 0; i < n; i++ {
		// first digit 1-9
		if i == 0 {
			c *= 9
		} else {
			c *= 10 - i
		}
	}

	return c + countNumbersWithUniqueDigits(n-1)
}

func main() {
	fmt.Printf("%v", countNumbersWithUniqueDigits(3))
}
