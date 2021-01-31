package solution

import "fmt"

func insert(intervals [][]int, newIntervals []int) [][]int {
	olflag := []int{}
	left := -1
	right := len(intervals)
	for i, row := range intervals {
		if overlap(row, newIntervals) {
			olflag = append(olflag, i)
			newIntervals[0] = min(row[0], newIntervals[0])
			newIntervals[1] = max(row[1], newIntervals[1])
		} else {
			if row[0] < newIntervals[0] {
				left = i
			}
			if row[0] > newIntervals[0] {
				right = i
				break
			}
		}
	}
	ans := [][]int{}
	fmt.Printf("append with %v, %v", left, right)

	ans = append(ans, intervals[0:left+1]...)
	ans = append(ans, newIntervals)
	ans = append(ans, intervals[right:]...)

	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func overlap(ar1, ar2 []int) bool {
	f1 := min(ar1[0], ar2[0])
	f2 := max(ar1[1], ar2[1])
	if f2-f1 <= ar1[1]+ar2[1]-ar1[0]-ar2[0] {
		return true
	}
	return false
}
