package HIndex2

func hIndex(citations []int) int {
	n := len(citations)
	left, right, pivot := 0, n-1, 0
	for left <= right {
		pivot = left + (right-left)/2
		if citations[pivot] == n-pivot {
			return n - pivot
		} else if citations[pivot] < n-pivot {
			left = pivot + 1
		} else {
			right = pivot - 1
		}

	}
	return n - left
}
