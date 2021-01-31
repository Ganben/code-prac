package solution

func maximalRectangle(matrix [][]byte) int {
	if matrix == nil {
		return 0
	}
	maxarea := 0
	dp := make([]int, len(matrix[0]))
	for _, row := range matrix {
		for j, nn := range row {
			if nn == '1' {
				dp[j] = dp[j] + 1
			} else {
				dp[j] = 0
			}

		}
		maxarea = max(maxarea, stackHelper(dp))
	}
	return maxarea
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func stackHelper(heights []int) int {
	stack := []int{-1}
	maxarea := 0
	for i, n := range heights {
		for stack[len(stack)-1] != -1 && heights[stack[len(stack)-1]] >= n {
			maxarea = max(maxarea, heights[stack[len(stack)-1]]*(len(heights)-stack[len(stack)-2]-1))
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	for stack[len(stack)-1] != -1 {
		maxarea = max(maxarea, heights[stack[len(stack)-1]]*(len(heights)-stack[len(stack)-2]-1))
		stack = stack[:len(stack)-1]
	}
	return maxarea
}
