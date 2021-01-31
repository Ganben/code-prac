package solution

import "fmt"

func numberOfSubarrays(nums []int, k int) int {
	// test odd num
	oddPos := []int{}
	total := len(nums)
	for i, n := range nums {
		if n%2 == 1 {
			oddPos = append(oddPos, i)
		}
	}
	if len(oddPos) < k {
		return 0
	}
	result := 0
	// more than one choice
	for i, n := range oddPos {
		// test break cond
		if i+k > len(oddPos) {
			fmt.Printf("end at %v\n", i)
			break
		}
		// test front choice
		c1 := 1
		if i == 0 {
			// all front can prepend
			c1 += n
		} else {
			c1 += n - oddPos[i-1] - 1
		}

		// test rear choice
		c2 := 1
		if i+k < len(oddPos) {
			c2 += oddPos[i+k] - oddPos[i+k-1] - 1
		} else if i+k == len(oddPos) {
			c2 += total - 1 - oddPos[i+k-1]
		}
		fmt.Printf("add to %v for %v * %v\n", result, c1, c2)
		result += c1 * c2
	}
	return result
}
