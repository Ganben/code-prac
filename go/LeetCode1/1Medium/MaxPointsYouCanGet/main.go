package main

import "fmt"

func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	windowSize := n - k
	sum := 0
	for _, pt := range cardPoints[:windowSize] {
		sum += pt
	}
	minSum := sum
	for i := windowSize; i < n; i++ {
		sum += cardPoints[i] - cardPoints[i-windowSize]
		minSum = min(minSum, sum)

	}
	total := 0
	for _, pt := range cardPoints {
		total += pt
	}
	return total - minSum
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 1}
	k := 3
	fmt.Printf("%v\n", maxScore(arr, k))
}
