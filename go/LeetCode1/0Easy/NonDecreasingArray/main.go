package main

import (
	"fmt"	"sort"
)


func check(nums []int) bool {
	for i:=0; i< len(nums)-1; i++ {
		x, y := nums[i], nums[i+1]
		if x > y {
			nums[i] = y
			if sort.IntsAreSorted(nums) {
				return true
			}
			nums[i] = x
			nums[i+1] = x
			return sort.IntsAreSorted(nums)
		}
	}
	return true
}

func checkPossibility(nums []int) bool {
	diff := make([]int, len(nums)-1)
	if len(nums) <= 1 {
		return false
	}
	change := []int{}
	for i := 1; i < len(nums); i++ {
		diff[i-1] = nums[i] - nums[i-1]
		if diff[i-1] < 0 {
			change = append(change, i)
		}
	}
	if len(change) == 0 {
		return true
	} else if len(change) > 2 {
		return false
	}

	if len(change) == 1 {
		if change[0] == 0 {
			return true
		} else if change[0] == len(nums)-1 {
			return true
		} else {
			i, j := change[0], change[0]-1
			if nums[i]-nums[j] < 1 {
				return true
			} else {
				return false
			}
		}
	} else if len(change) == 2 {
		if change[0]+1 != change[1] {
			return false
		}
		i, k := change[0]-1, change[1]
		if nums[i]+1 < nums[k] {
			return true
		} else {
			return false
		}
	}
	return false
}

func main() {
	arr := []int{4, 2, 3}
	fmt.Printf("%v\n", checkPossibility(arr))

}
