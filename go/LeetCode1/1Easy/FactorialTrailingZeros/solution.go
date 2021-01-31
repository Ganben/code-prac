package FactorialTrailingZeros

func trailingZeros(n int) int {
	zeroCount := 0
	// currentMul := 5
	for n > 0 {
		n /= 5
		zeroCount += n
	}
	return zeroCount
}
