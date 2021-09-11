package nonnegintwithoutconsecones

import "testing"

func TestSolve(t *testing.T) {
	res := findIntegers(5)
	if res != 5 {
		t.Log(res)
		t.FailNow()
	}

}
