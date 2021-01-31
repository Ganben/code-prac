package main

import "fmt"

func countBalls(lowLimit int, highLimit int) int {
	countBox := map[int]int{}
	for i := lowLimit; i <= highLimit; i++ {
		p := sumNum(i)
		_, exist := countBox[p]
		if exist {
			countBox[p]++
		} else {
			countBox[p] = 1
		}
	}
	//
	maxN := 0
	for _, v := range countBox {
		if v > maxN {
			maxN = v
		}
	}
	return maxN
}

func sumNum(n int) int {
	arr := []int{}
	for n > 0 {
		r := n % 10
		arr = append(arr, r)
		n /= 10
	}
	ans := 0
	for _, v := range arr {
		ans += v
	}
	return ans
}

func main() {
	l, h := 1, 10
	fmt.Printf("%v\n", countBalls(l, h))
	l, h = 19, 28
	fmt.Printf("%v\n", countBalls(l, h))
}
