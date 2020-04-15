package solution

func merge(intervals [][]int) [][]int {
	var ans [][]int
	sorted := mergeSort(intervals)
	flag := false
	for i, row := range sorted {
		if i == 0 {
			ans = append(ans, row)
			continue
		}
		flag = false
		for _, el := range ans {
			if overlap(el, row) {
				el[0] = min(el[0], row[0])
				el[1] = max(el[1], row[1])
				flag = true
				continue
			}
		}
		if !flag {
			ans = append(ans, row)
		}
	}
	return ans
}

func mergeSort(arr [][]int) [][]int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	middle := int(length / 2)
	left := arr[0:middle]
	right := arr[middle:]
	return mergeOrigin(mergeSort(left), mergeSort(right))
}

func mergeOrigin(left, right [][]int) [][]int {
	var result [][]int
	// main loop
	for len(left) != 0 && len(right) != 0 {
		if left[0][0] <= right[0][0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}
	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}
	return result
}

func overlap(x []int, ar []int) bool {
	// left distance f1 >0 or f1<0 while true
	f1 := min(x[0], ar[0])
	// right distance f2 >0 or f2 < 0 while true
	f2 := max(x[1], ar[1])
	// fmt.Printf("overlaped test %v,%v=%v\n", x, ar, !(f1 || f2))
	if f2-f1 <= x[1]+ar[1]-x[0]-ar[0] {
		return true
	}
	return false
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
