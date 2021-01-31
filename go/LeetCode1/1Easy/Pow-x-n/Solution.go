package Solution

func myPow(x float64, n int) float64 {
	//
	if n == 0 {
		return 1
	}
	//
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	//
	if n%2 == 1 {
		return x * myPow(x, n-1)
	}
	//
	if n%2 == 0 {
		return myPow(x*x, n/2)
	}
	return 0.0
}
