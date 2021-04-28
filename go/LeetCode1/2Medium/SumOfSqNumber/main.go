package main

import "fmt"

var squared map[int]int

func judgeSquareSum(c int) bool {
	if c == 1 || c == 0 {
		return true
	}

	// squared[1] = 1
	// f := true
	for i := 1; i < c; i++ {
		squared[i*i] = i
		if i*i >= c {
			break
		}
	}
	if _, exist := squared[c]; exist {
		return true
	}
	//
	for k, _ := range squared {
		if _, exist := squared[c-k]; exist {
			return true
		}
	}
	return false
}

func sqrtLow(c int) int {
	l := 0
	for i := 1; i < c; i++ {
		if c>>i == 0 {
			l = i
			break
		}
	}
	//
	ans := 1
	return ans << (l / 2)
}

func main() {
	c := 5
	fmt.Printf("%v\n", judgeSquareSum(c))
}
