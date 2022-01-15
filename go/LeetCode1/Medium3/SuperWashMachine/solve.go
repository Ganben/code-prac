package superwashmachine

func findMinMoves(machines []int) int {
	ans := 0
	tot := 0
	for _, v := range machines {
		tot += v

	}
	n := len(machines)
	if tot%n > 0 {
		return -1
	}
	avg := tot / n
	sum := 0
	for _, num := range machines {
		num -= avg
		sum += num
		ans = max(ans, max(abs(sum), num))
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
