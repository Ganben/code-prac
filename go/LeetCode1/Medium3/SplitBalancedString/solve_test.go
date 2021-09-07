package splitbalancedstring

import "testing"

func TestSolve(t *testing.T) {
	s := "RLRRLLRLRL"
	res := balancedStringSplit(s)
	if res != 4 {
		t.Log(res)
		t.FailNow()
	}
}
