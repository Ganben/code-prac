package Solution

var qstate [][]int

func updateMatrix(matrix [][]int) [][]int {

	lm := len(matrix)
	if lm == 0 {
		return [][]int{}
	}
	ln := len(matrix[0])
	qstate = [][]int{}
	for i := 0; i < lm; i++ {
		row := make([]int, ln)
		for j := 0; j < ln; j++ {
			row[j] = 2 * (lm + ln)
		}
		qstate = append(qstate, row)
	}

	for i := 0; i < lm; i++ {
		for j := 0; j < ln; j++ {
			if matrix[i][j] == 0 {
				qstate[i][j] = 0
			}
		}
	}

	// direction 1
	for i := 0; i < lm; i++ {
		for j := 0; j < ln; j++ {
			window(i, j)
		}
	}
	// direction 2
	for j := ln - 1; j >= 0; j-- {
		for i := lm - 1; i >= 0; i-- {
			window(i, j)
		}
	}
	// direction 3
	for i := 0; i < lm; i++ {
		for j := ln - 1; j >= 0; j-- {
			window(i, j)
		}
	}

	// direction 4
	for i := lm - 1; i >= 0; i-- {
		for j := 0; j < ln; j++ {
			window(i, j)
		}
	}

	// find stack = 0 and update nearby window
	return qstate
}

func window(i, j int) {
	// fmt.Printf("visit %v,%v\n", i, j)
	nears := []int{}

	if i-1 >= 0 {
		nears = append(nears, qstate[i-1][j])
	}
	if j-1 >= 0 {
		nears = append(nears, qstate[i][j-1])
	}
	if i+1 < len(qstate) {
		nears = append(nears, qstate[i+1][j])
	}
	if j+1 < len(qstate[0]) {
		nears = append(nears, qstate[i][j+1])
	}
	m := minarray(nears)
	qstate[i][j] = min2(qstate[i][j], m+1)

}

func minarray(ar []int) int {
	switch len(ar) {
	case 1:
		return ar[0]
	case 2:
		return min2(ar[0], ar[1])
	case 3:
		return min3(ar[0], ar[1], ar[2])
	case 4:
		return min4(ar[0], ar[1], ar[2], ar[3])
	}
	return 0
}

func min4(x, y, z, t int) int {
	return min2(min2(x, y), min2(z, t))
}

func min3(x, y, z int) int {
	return min2(x, min2(y, z))
}

func min2(x, y int) int {
	if x > y {
		return y
	}
	return x
}
