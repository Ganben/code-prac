package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	ans := []string{}
	// nextI := 1
	state := []int{}
	for i, n := range nums {
		if i == len(nums)-1 {
			// check state
			if len(state) != 0 {
				s := strconv.Itoa(state[0]) + "->" + strconv.Itoa(n)
				state = []int{}
				ans = append(ans, s)
			} else {
				// new range
				ans = append(ans, strconv.Itoa(n))
			}
			break
		}

		if len(state) == 0 {
			// new range
			if n+1 == nums[i+1] {
				state = append(state, n)
			} else {
				ans = append(ans, strconv.Itoa(n))
			}

		} else {
			// extend range
			if n+1 == nums[i+1] {
				state = append(state, n)
			} else {
				s := strconv.Itoa(state[0]) + "->" + strconv.Itoa(n)
				state = []int{}
				ans = append(ans, s)

			}
			// skip

		}
	}
	return ans
}

func main() {
	intsarr := []int{0, 1, 2, 4, 5, 7}
	fmt.Printf("%v", summaryRanges(intsarr))
}
