package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	hmap map[int]int
	nums []int
}

func Constructor(nums []int) Solution {
	hmap := map[int]int{}

	for i := 0; i < len(nums); i++ {
		hmap[i] = i
	}
	return Solution{hmap, nums}

}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	for i := 0; i < len(this.nums); i++ {
		this.hmap[i] = i
	}
	return this.nums
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	arr := []int{}
	for i := 0; i < len(this.nums); i++ {
		arr = append(arr, i)
	}
	for i := 0; i < len(this.nums); i++ {
		b := rand.Intn(len(this.nums) - i)
		this.hmap[i] = arr[b]
		arr = append(arr[:b], arr[b+1:]...)
	}
	ans := []int{}
	for i := 0; i < len(this.nums); i++ {
		ans = append(ans, this.nums[this.hmap[i]])
	}
	return ans
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

func main() {
	nums := []int{1, 2, 3}
	obj := Constructor(nums)
	param_1 := obj.Reset()
	fmt.Printf("%v\n", param_1)
	param_2 := obj.Shuffle()
	fmt.Printf("%v\n", param_2)
}
