package NumberOfDigitOne

func countDigitOne(n int) int {
	countr := 0
	for i := 1; i <= n; i *= 10 {
		divider := i * 10
		countr += (n/divider)*i + min(max(n%divider-i+1, 0), i)
	}
	return countr
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
