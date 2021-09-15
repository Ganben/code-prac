package randompointinnorect

import "testing"

func TestSolve(t *testing.T) {
	rects := [][]int{[]int{1, 1, 5, 5}}
	obj := Constructor(rects)
	params1 := obj.Pick()
	if params1[0] >= 5 || params1[0] <= 1 {
		t.Log(params1)
		t.FailNow()
	}
}
