package main

import (
	"fmt"
	"math"
	"sort"
)

type KthLargest struct {
	arr []int
	kth int
}

func Constructor(k int, nums []int) KthLargest {
	arrs := make([]int, k)
	sort.Ints(nums) //very slow
	fmt.Printf("%v\n", nums)
	for i := 0; i < k; i++ {
		if i > len(nums)-1 {
			arrs[i] = 1 - math.MaxInt32
		} else {
			arrs[i] = nums[len(nums)-1-i]
		}
	}
	fmt.Printf("%v\n", arrs)
	return KthLargest{arrs, k}
}

func (this *KthLargest) Add(val int) int {
	if len(this.arr) == 0 {
		this.arr = append(this.arr, val)
		return val
	}
	if len(this.arr) == 1 {
		if val > this.arr[0] {
			this.arr[0] = val
			return val
		}
	}
	for i := len(this.arr) - 1; i > 0; i-- {
		if this.arr[i] < val && this.arr[i-1] >= val {
			// insert between
			head := []int{}
			head = append(head, this.arr[:i]...)
			head = append(head, val)
			head = append(head, this.arr[i:len(this.arr)-1]...)
			this.arr = head
			fmt.Printf("insert %v\n", head)
			break
		} else if i == 1 && this.arr[i-1] < val {
			// insert head
			head := []int{val}
			head = append(head, this.arr[:len(this.arr)-1]...)
			this.arr = head
			fmt.Printf("head %v\n", head)
		}
	}
	return this.arr[len(this.arr)-1]
}

func main() {
	k := 3
	nums := []int{4, 5, 8, 2}
	obj := Constructor(k, nums)
	// fmt.Printf("%v\n", nums)
	varArr := []int{3, 5, 10, 9, 4}
	for _, val := range varArr {
		fmt.Printf("%v\n", obj.Add(val))
	}

}
