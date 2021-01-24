package main

import "fmt"

var combs [][]int
var choices []int

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, n := range nums {
			if n <= i {
				dp[i] += dp[i-n]
			}
		}
	}
	return dp[target]
}

func combinationSum(nums []int, target int) int {
	combs = [][]int{}
	choices = nums
	cur := []int{}
	dfs(cur, target)
	fmt.Printf("%v", combs)
	return len(combs)
}

func dfs(curr []int, target int) {

	p := sum(curr)
	if p == target {
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		combs = append(combs, tmp)
		// fmt.Printf("dfs call %v", curr)
		return
	} else if p < target {
		for _, v := range choices {
			newCur := append(curr, v)
			dfs(newCur, target)
		}
	}
	return
}

func sum(arr []int) int {
	ans := 0
	for _, v := range arr {
		ans += v
	}
	return ans
}

func main() {
	nums := []int{4, 2, 1}
	target := 32
	fmt.Printf("%v", combinationSum4(nums, target))
}
