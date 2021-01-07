package main

import "fmt"

func lengthOfLIS(nums []int) int {
	lis := []int{}

	for i, v := range nums {
		if i == 0 {
			lis = append(lis, nums[0])
			continue
		}
		//
		if lis[len(lis)-1] < v {
			lis = append(lis, v)
			fmt.Printf("%v\n", lis)
		} else {
			// bin search
			l := 0
			r := len(lis) - 1
			pos := r
			for l <= r {
				mid := (l + r) / 2
				if lis[mid] >= v {
					pos = mid
					r = mid - 1
				} else {
					l = mid + 1
				}

			}
			// lis[pos]
			lis[pos] = v
			fmt.Printf("%v\n", lis)
		}
	}
	return len(lis)
}

func main() {
	inputArr := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Printf("%v", lengthOfLIS(inputArr))

}
