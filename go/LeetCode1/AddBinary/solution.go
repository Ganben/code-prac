package solution

func addBinary(a, b string) []int {
	rarr := make([]int, 1+max(len(a), len(b)))
	offset := 0
	for i := len(rarr) - 1; i > 0; i-- {
		//
		if offset == min(len(a), len(b)) {
			break
		}
		r := binadd(rune(a[len(a)-1-offset]), rune(b[len(a)-1-offset]))
		if r == 2 {
			if rarr[i] == 1 {
				rarr[i] = 1
				rarr[i-1] = 1
			} else {
				rarr[i] = 0
				rarr[i-1] = 1
			}

		} else if r == 1 {
			if rarr[i] == 1 {
				rarr[i-1] = 1
				rarr[i] = 0
			} else {
				rarr[i] = 1
			}

		}
		offset++
	}
	return rarr

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func binadd(a, b rune) int {
	va := 0
	vb := 0
	if a == '1' {
		va = 1
	}
	if b == '1' {
		vb = 1
	}
	return va + vb
}
