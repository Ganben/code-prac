package main

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	p1, p2, p3 := math.MaxInt32, 0, -1*(math.MaxInt32-1)
	n1, n2, n3 := 0, 0, 0
	for i := 0; i <= len(nums)-3; i++ {
		if p1 > nums[i] {
			p1 = nums[i]
			n1 = i
		}
		if nums[i+1] > p1 {
			p2 = nums[i+1]
			n2 = i + 1
		}
		if p3 < nums[i+2] {
			p3 = nums[i+2]
			n3 = i + 2
		}
	}
	if p1 < p2 && p2 < p3 {
		if n1 < n2 && n2 < n3 {
			return true
		}
	}
	return false
}

func increasingTripletSimple(nums []int) bool {
	for i := 0; i <= len(nums)-3; i++ {
		if nums[i] < nums[i+1] && nums[i+1] < nums[i+2] {
			return true
		}
	}
	return false
}

func main() {
	i := []int{1, 2, 3, 4, 5}
	r1 := increasingTriplet(i)

	i2 := []int{1, 2, -10, -8, -7}
	r2 := increasingTriplet(i2)
	fmt.Printf("v%,v%", r1, r2)
}
