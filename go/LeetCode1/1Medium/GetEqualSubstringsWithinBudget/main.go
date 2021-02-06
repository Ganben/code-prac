package main

import "fmt"

func equalSubstring(s string, t string, maxCost int) int {
	maxLen := 0
	n := len(s)
	diff := make([]int, n)
	for i, ch := range s {
		diff[i] = abs(int(ch) - int(t[i]))
	}
	sum, start := 0, 0
	for end, d := range diff {
		sum += d
		for sum > maxCost {
			sum -= diff[start]
			start++
		}
		maxLen = max(maxLen, end-start+1)
	}
	return maxLen
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	s := "abcd"
	t := "bcdf"
	cost := 3
	fmt.Printf("%v\n", equalSubstring(s, t, cost))
}
