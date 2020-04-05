package Solution

import (
	"fmt"
)

func trap(height []int) int {
	v := 0
	leftWall := -1
	rightWall := -1
	for i, n := range height {
		if i > 0 {
			// if n < height[i-1] && leftWall < 0 {
			// 	leftWall = i
			// 	fmt.Println("got left", i)
			// } else if n > height[i-1] && leftWall > 0 {
			// 	rightWall = i
			// 	fmt.Println("got right", i)
			// } else if n < height[i-1] && leftWall >= 0 && rightWall > 0 {
			// 	v += volumn(height[leftWall:rightWall])
			// 	leftWall = -1
			// 	rightWall = -1
			// 	fmt.Println("call v",i)
			// }
			if n < height[i-1] {
				if leftWall < 0 && rightWall < 0 {
					leftWall = i
					fmt.Println("got left", i)
				} else if leftWall > 0 && rightWall < 0 {
					continue
				} else if leftWall < 0 && rightWall > 0 {
					leftWall = i
					rightWall = -1
					fmt.Println("got left", i)
				} else if leftWall > 0 && rightWall > 0 {
					v += volumn(height[leftWall-1:rightWall+1])
					fmt.Println("call v and update ", i, leftWall, rightWall)
					leftWall = i
					rightWall = -1
				}
			} else if n > height[i-1] {
				if leftWall < 0 && rightWall < 0 {
					continue
				} else if leftWall > 0 && rightWall < 0 {
					rightWall = i
					fmt.Println("got right", i)
				} else if leftWall < 0 && rightWall > 0 {
					continue
				} else if leftWall > 0 && rightWall > 0 {
					rightWall = i
					fmt.Println("update right", i)
				}
			}
		} else {
			continue
		}

	}
	return v
}

func volumn(height []int) int {
	d := len(height)
	h := min(height[0], height[d-1])
	v := 0
	for _, n := range height {
		if n < h {
			v += h - n
			fmt.Println("set add ", h - n, h)
		}
	}
	fmt.Println("add v ", v)
	return v
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}