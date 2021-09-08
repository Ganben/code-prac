package ipo

import "testing"

func TestSolve(t *testing.T) {
	k := 2
	w := 0
	profits := []int{1, 2, 3}
	capital := []int{0, 1, 1}
	res := findMaximizedCapital(k, w, profits, capital)
	if res != 4 {
		t.Log(res)
		t.FailNow()
	}
}
