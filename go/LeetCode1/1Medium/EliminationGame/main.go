package main

import "fmt"

func lastRemaining(n int) int {
	if n == 1 {
		return 1
	}
	return 2 * (n/2 - lastRemaining(n/2) + 1)
}

func main() {
	fmt.Printf("%v", lastRemaining(9))
}
