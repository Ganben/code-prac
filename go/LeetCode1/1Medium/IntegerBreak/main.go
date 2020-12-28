package main

import (
	"fmt"
	"math"
)

func integerBreak(n int) int {
	maxMultip := 0
	resArr := []int{0, 0, 0}
	curs := []int{}
	for i := 1; i < n; i++ {
		if n == 1 {
			return 1
		}
		r := 1
		times := 1
		if i == 1 {
			maxMultip = 1
		} else {
			r = n % i
			times = int(n / i)
			if times > 1 && r != 0 {
				v1 := IPow(i, times-1) * (i + r)
				v2 := IPow(i, times) * r
				if v1 > v2 {
					maxMultip = v1
				} else {
					maxMultip = v2
				}
			} else if times == 1 {
				if r == 0 {
					maxMultip = i
				} else {
					maxMultip = i * r
				}
			} else if times > 1 && r == 0 {
				maxMultip = IPow(i, times)
			}

		}

		if i == 1 {
			resArr[0] = maxMultip
		}
		if i == 2 {
			// resArr[0] = resArr[1]
			resArr[1] = maxMultip
		}
		if i == 3 {
			// resArr[0] = resArr[1]
			// resArr[1] = resArr[2]
			resArr[2] = maxMultip
		}
		if i > 3 {
			resArr[0] = resArr[1]
			resArr[1] = resArr[2]
			resArr[2] = maxMultip
		}
		//
		if i >= 3 && resArr[0] <= resArr[1] && resArr[1] >= resArr[2] {
			fmt.Printf("%v", resArr)
			curs = append(curs, resArr[1])
		}

	}
	fmt.Print(resArr)
	if maxMultip > RowMax(curs) {
		return maxMultip
	} else {
		return RowMax(curs)
	}

}

func RowMax(arr []int) int {
	ans := 0
	for _, v := range arr {
		if v > ans {
			ans = v
		}
	}
	return ans
}

func IPow(x, y int) int {
	xx, yy := float64(x), float64(y)
	return int(math.Pow(xx, yy))
}

func main() {
	fmt.Printf("%v", integerBreak(8))
}
