package main

import (
	"fmt"
	"math"
)

func leastInterval(tasks []byte, n int) int {
	minTime := 0
	cnt := map[byte]int{}
	for _, t := range tasks {
		cnt[t]++
	}
	nextValid := make([]int, 0, len(cnt))
	rest := make([]int, 0, len(cnt))
	for _, c := range cnt {
		nextValid = append(nextValid, 1)
		rest = append(rest, c)
	}

	for range tasks {
		minTime++
		minNextValid := math.MaxInt64
		for i, r := range rest {
			if r > 0 && nextValid[i] < minNextValid {
				minNextValid = nextValid[i]
			}
		}
		if minNextValid > minTime {
			minTime = minNextValid
		}
		best := -1
		for i, r := range rest {
			if r > 0 && nextValid[i] <= minTime && (best == -1 || r > rest[best]) {
				best = i
			}
		}
		nextValid[best] = minTime + n + 1
		rest[best]--

	}
	return minTime
}

func main() {
	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	n := 2
	r := leastInterval(tasks, n)
	fmt.Printf("%v", r)
}
