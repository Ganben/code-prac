package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Solution struct {
	r    *rand.Rand
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{r: rand.New(rand.NewSource(time.Now().Unix())), nums: nums}
}

func (this *Solution) Pick(target int) int {
	k := 1
	var res int
	for i := 0; i < len(this.nums); i++ {
		if this.nums[i] == target {
			if this.r.Intn(k) == 0 {
				res = i
			}
			k++
		}
	}
	return res
}

func main() {
	nums := []int{1, 2, 3, 3, 3}
	obj := Constructor(nums)
	target := 3
	param_1 := obj.Pick(target)
	fmt.Printf("%v\n", param_1)
}
