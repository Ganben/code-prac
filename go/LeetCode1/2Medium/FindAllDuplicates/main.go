package main

import "fmt"

func findDuplicates(nums []int) []int {
	// xs := make([]bool, len(nums))
	res := []int{}
	for _, v := range nums {
		id := 0
		if v < 0 {
			id = -v - 1
		} else {
			id = v - 1
		}
		if nums[id] < 0 {
			res = append(res, id+1)
		} else {
			nums[id] *= -1
		}
	}
	return res

}

func main() {
	i := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Printf("%v\n", findDuplicates(i))
}
