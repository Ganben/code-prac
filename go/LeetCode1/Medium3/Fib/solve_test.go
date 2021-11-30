package fib

import "testing"

func TestSolve(t *testing.T) {
	for i := 0; i < 10; i++ {
		a := fib(i)
		println(a)
	}
}
