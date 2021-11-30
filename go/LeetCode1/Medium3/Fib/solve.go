package fib

func fib(n int) int {
	if n == 1 {
		return 1
	}
	if n == 0 {
		return 0
	}

	a1 := 0
	a2 := 1
	for i := 2; i < n; i++ {
		tmp := a1 + a2
		a1 = a2
		a2 = tmp
	}
	return a2 + a1
}
