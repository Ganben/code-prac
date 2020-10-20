package PerfectSquares

var squareNums map[int]int

func numSquares(n int) int {
	squareNums = map[int]int{}
	for i := 1; i*i <= n; i++ {
		squareNums[i*i] = 1
	}
	c := 1
	for count := 1; count <= n; count++ {
		if isDividedBy(n, count) {
			return count
		}
		c = count
	}
	return c
}

func isDividedBy(n, count int) bool {
	if count == 1 {
		if _, ok := squareNums[n]; ok {
			return true
		}
	}

	for _, k := range squareNums {
		if isDividedBy(n-k, count-1) {
			return true
		}
	}
	return false
}
