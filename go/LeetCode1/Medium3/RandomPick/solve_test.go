package randompick

import "testing"

func TestSolve2(t *testing.T) {
	t.Log("start test")
	w := []int{1}
	obj := Constructor(w)
	param_1 := obj.PickIndex()
	if param_1 != 0 {
		t.Log(param_1)
		t.FailNow()
	}
}
