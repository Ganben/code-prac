package base7

import "testing"

func TestSolve(t *testing.T) {
	n := -7
	res := convertToBase7(n)
	if res != "-10" {
		t.Log(res)
		t.FailNow()
	}
}
