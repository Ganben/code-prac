package BitwiseAndRange

func rangeBitwiseAnd(m, n int) int {
	for m < n {
		n = n & (n - 1)
	}
	return m & n
}
