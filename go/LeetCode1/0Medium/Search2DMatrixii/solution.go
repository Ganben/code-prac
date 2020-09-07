package Search2DMatrixii

func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 {
		return false
	}
	shorterDim := min(len(matrix), len(matrix[0]))
	for i := 0; i < shorterDim; i++ {
		verticalFound := binarySearch(matrix, target, i, true)
		horizontalFound := binarySearch(matrix, target, i, false)
		if verticalFound || horizontalFound {
			return true
		}
	}
	return false
}

func binarySearch(matrix [][]int, target, start int, vertical bool) bool {
	lo := start
	var hi int
	if vertical {
		hi = len(matrix[0]) - 1
	} else {
		hi = len(matrix) - 1
	}
	for hi >= lo {
		mid := (lo + hi) / 2
		if vertical {
			if matrix[start][mid] < target {
				lo = mid + 1
			} else if matrix[start][mid] > target {
				hi = mid - 1
			} else {
				return true
			}
		} else {
			if matrix[mid][start] < target {
				lo = mid + 1
			} else if matrix[mid][start] > target {
				hi = mid - 1
			} else {
				return true
			}
		}
	}
	return false
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
