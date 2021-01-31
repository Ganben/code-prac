package solution

func mySqrt(x int) int {
	if x < 1 {
		return 0
	}
	if x == 1 {
		return 1
	}
	if x == 2 {
		return 1
	}
	prev := false
	for i := mySqrt(int(x / 2)); i <= int(x*3/4); i++ {
		if i*i < x {
			prev = true
		}
		if i*i == x {
			return i
		}
		if i*i > x {
			if prev {
				return i - 1
			}
			break
		}
	}
	return -1
}
