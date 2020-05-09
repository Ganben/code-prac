package Search2DMatrix

func searchMatrix(matrix [][]int, target int) bool {
	pos := -1
	left, right := 0, len(matrix)
	//row bin search
Loop1:
	for {
		if left >= right {
			break Loop1
		}
		mid := int((left + right) / 2)
		if mid == 0 {
			pos = mid
			break Loop1
		}
		switch true {
		case matrix[mid][0] < target:
			if left == mid {
				pos = mid
				break Loop1
			}
			left = mid
			if mid+1 < len(matrix) && matrix[mid+1][0] > target {
				pos = mid
				break Loop1
			}
		case matrix[mid][0] > target:
			if right == mid {
				pos = mid
				break Loop1
			}
			right = mid
			if mid-1 >= 0 && matrix[mid-1][0] < target {
				pos = mid - 1
				break Loop1
			}
		case matrix[mid][0] == target:
			pos = mid
			return true
		}
	}
	if pos == -1 {
		return false
	}
	// col bin search
	rpos := -1
	left, right = 0, len(matrix[pos])
Loop2:
	for {
		if left >= right {
			break Loop2
		}
		mid := int((left + right) / 2)
		switch true {
		case matrix[pos][mid] < target:
			if left == mid {
				return false
			}
			left = mid
			if left+1 < len(matrix[pos]) && matrix[pos][left+1] > target {
				return false
			}

		case matrix[pos][mid] > target:
			if right == mid {
				return false
			}
			right = mid
			if right-1 >= 0 && matrix[pos][right-1] < target {
				return false
			}

		case matrix[pos][mid] == target:
			rpos = mid
			break Loop2
		}

	}
	if rpos == -1 {
		return false
	}
	return true
}
