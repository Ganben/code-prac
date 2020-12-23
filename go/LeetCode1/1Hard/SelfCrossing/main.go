package main

import "fmt"

func isSelfCrossing(x []int) bool {
	if len(x) < 4 {
		return false
	}
	for i := 3; i < len(x); i++ {
		if i >= 3 {
			if isSelfCrossing4Node(x[i-3 : i+1]) {
				return true
			}
		}
		if i >= 4 {
			if isSelfCrossing5Node(x[i-4 : i+1]) {
				return true
			}
		}
		if i >= 5 {
			if isSelfCrossing6Node(x[i-5 : i+1]) {
				return true
			}
		}
	}
	return false
}

func isSelfCrossing4Node(x []int) bool {
	fmt.Printf("test %v", x)
	if x[0] < 0 {
		return false
	}
	dx, dy := 0, 0
	// lx, ly := 0,0
	dy = x[0] - x[2]
	dx = x[3] - x[1]
	if dy >= 0 && dx >= 0 {
		return true
	}
	return false
}

func isSelfCrossing5Node(x []int) bool {
	fmt.Printf("test %v", x)
	if x[0] < 0 {
		return false
	}
	if x[1] == x[3] && x[4]+x[0] >= x[2] {
		return true
	}
	return false
}

func isSelfCrossing6Node(x []int) bool {
	fmt.Printf("test %v", x)
	if x[0] < 0 {
		return false
	}
	dx, dy := 0, 0
	dx = x[3] - x[1] - x[5]
	dy = x[0] - x[2] + x[4]
	if dx <= 0 && dy >= 0 && x[1] < x[3] && x[2] > x[4] {
		return true
	}
	return false
}

func main() {
	x := []int{2, 2, 3, 3, 2, 2}
	fmt.Printf("%v", isSelfCrossing(x))
}
