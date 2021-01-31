package Solution

func maxArea(height []int) int {
	max, l, r := 0, 0, len(height) - 1
	for l < r {
		max = Max(max, Min(height[l], height[r]) * (r - l))
		if height[l] < height[r] {
			l ++
		} else {
			r --
		}
	}
	return max
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}