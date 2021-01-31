package MaxPointsOnALine

var pointsG [][]int
var n int
var lines map[float64]int
var horisontalLines int

func addLine(i, j, count, duplicate int) (a, b int) {
	x1 := pointsG[i][0]
	y1 := pointsG[i][1]
	x2 := pointsG[j][0]
	y2 := pointsG[j][1]

	if x1 == x2 && y1 == y2 {
		duplicate++
	} else if y1 == y2 {
		horisontalLines++
		count = max(horisontalLines, count)
	} else {
		slope := float64(1.0*(x1-x2)/(y1-y2) + 0.0)
		if v, ok := lines[slope]; ok {
			lines[slope] = v + 1
		} else {
			lines[slope] = 2
		}
		count = max(lines[slope], count)
	}
	a = count
	b = duplicate
	return a, b
}

func maxPointsOnLineContainingPoint(i int) int {

}

func maxPoints(points [][]int) int {
	pointsG = points
	n := len(points)
	if n < 3 {
		return n
	}
	maxCount := 1
	for i := 0; i < n-1; i++ {
		maxCount = max(maxPointsOnLineContainingPoint(i), maxCount)
	}
	return maxCount
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
