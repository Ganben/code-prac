package main

import "fmt"

func findLengthOfLCIS(nums []int) int {
	// state := true
	if len(nums) < 1 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	ans := 1
	ansArr := []int{}
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			// state = true
			ans += 1
			fmt.Printf("%v", ans)
		} else {
			// state = false
			ansArr = append(ansArr, ans)
			ans = 1
		}
	}
	for _, v := range ansArr {
		if v > ans {
			ans = v
		}
	}
	return ans

}

func main() {
	nums := []int{1, 3, 5, 7}
	fmt.Printf("%v", findLengthOfLCIS(nums))
}
