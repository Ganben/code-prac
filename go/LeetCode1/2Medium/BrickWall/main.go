package main

import "fmt"

func leastBricks(wall [][]int) int {
	cnt := map[int]int{}
	for _, widths := range wall {
		sum := 0
		for _, width := range widths[:len(widths)-1] {
			sum += width
			cnt[sum]++
		}
	}
	maxCnt := 0
	for _, c := range cnt {
		if c > maxCnt {
			maxCnt = c
		}
	}
	return len(wall) - maxCnt
}

func main() {
	i := [][]int{
		{1, 2, 2, 1},
		{3, 1, 2},
		{1, 3, 2},
		{2, 4},
		{3, 1, 2},
		{1, 3, 1, 1},
	}
	fmt.Printf("%v\n", leastBricks(i))
}
