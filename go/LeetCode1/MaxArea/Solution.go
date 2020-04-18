package solution

func maxArea(height []int) int {
	// search two times and fill each value for maximum contains
	var eachMax = make([]int, len(height))
	for i, n := range height {
		for j, m := range height {
			eachMax[i] = max(eachMax[i], area(i, n, j, m))
		}
	}
	res := 0
	for _, n := range eachMax {
		if res < n {
			res = n
		}
	}
	return res
}

func area(i, n, j, m int) int {
	return abs(i, j) * min(n, m)
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
