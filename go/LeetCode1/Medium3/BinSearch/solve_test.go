package binsearch

import "testing"

func TestSolve(t *testing.T) {
	nArr := []int{-1, 0, 3, 5, 9, 12}
	targ := 9
	res := search(nArr, targ)
	if res != 4 {
		t.Log(res)
		t.FailNow()
	}
}
