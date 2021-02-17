package main

import "fmt"

func readBinaryWatch(num int) []string {
	res := []string{}
	backtrack(0, 10, num, []int{}, &res)
	return res
}

func backtrack(start, cap, target int, path []int, res *[]string) {
	if len(path) == target {
		min, hour := 0, 0
		for _, v := range path {
			if v >= 4 {
				min += 1 << (v - 4)
			} else {
				hour += 1 << v
			}
		}
		if min < 60 && hour < 12 {
			*res = append(*res, fmt.Sprintf("%d:%02d", hour, min))
		}
		path = []int{}
		return
	}
	//
	for i := start; i < cap; i++ {
		path = append(path, i)
		backtrack(i+1, cap, target, path, res)
		path = path[:len(path)-1]
	}
}

func main() {
	fmt.Printf("%v\n", readBinaryWatch(1))
}
