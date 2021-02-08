package main

import "fmt"

func maxTurbulenceSize(arr []int) int {
	if len(arr) < 2 {
		return len(arr)
	}
	n := len(arr)
	ans := 1
	left, right := 0, 0
	for right < n-1 {
		if left == right {
			if arr[left] == arr[left+1] {
				left++
			}
			right++
		} else {
			if arr[right-1] < arr[right] && arr[right] > arr[right+1] {
				right++
			} else if arr[right-1] > arr[right] && arr[right] < arr[right+1] {
				right++
			} else {
				left = right
			}
		}
	}
	ans = max(ans, right-left+1)
	return ans
}

func max2(arr []int) int {
	n := len(arr)
    dp := make([][2]int, n)
    dp[0] = [2]int{1, 1}
    for i := 1; i < n; i++ {
        dp[i] = [2]int{1, 1}
        if arr[i-1] > arr[i] {
            dp[i][0] = dp[i-1][1] + 1
        } else if arr[i-1] < arr[i] {
            dp[i][1] = dp[i-1][0] + 1
        }
    }

    ans := 1
    for i := 0; i < n; i++ {
        ans = max(ans, dp[i][0])
        ans = max(ans, dp[i][1])
    }
    return ans

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/longest-turbulent-subarray/solution/zui-chang-tuan-liu-zi-shu-zu-by-leetcode-t4d8/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	arr := []int{9, 4, 2, 10, 7, 8, 8, 1, 9}
	fmt.Printf("%v\n", maxTurbulenceSize(arr))
}
