package Powerof3

func isPowerOfThree(n int) bool {
	if n > 3 {
		floatRes := float64(n) / 3
		intRes := n / 3
		if int(floatRes*10) > intRes*10 {
			return false
		} else {
			return isPowerOfThree(n / 3)
		}

	} else if n == 3 {
		return true
	} else if n == 1 {
		return true
	} else {
		return false
	}
}

func isPowerOfThree3(n int) bool {
	return n > 0 && power(3, 19)%n == 0
}

//1162261467

func power(x, y int) int {
	ans := 1
	for i := 0; i < y; i++ {
		ans *= x
	}
	return ans
}
