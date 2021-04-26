package main

import (
	"fmt"
	"sort"
)

func shipWithinDays(weights []int, D int) int {
	left, right := 0, 0
	for _, w := range weights {
		if w > left {
			left = w
		}
		right += w
	}
	return left + sort.Search(right-left, func(x int) bool {
		x += left
		day := 1
		sum := 0
		for _, w := range weights {
			if sum+w > x {
				day++
				sum = 0
			}
			sum += w
		}
		return day <= D
	})
}

func main() {
	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	D := 5
	fmt.Printf("%v\n", shipWithinDays(weights, D))
}
