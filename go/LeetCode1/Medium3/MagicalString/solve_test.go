package magicalstring

import "testing"

func TestSolve(t *testing.T) {
	n := 6
	r := 3
	if magicalString(n) != r {
		t.FailNow()
	}
	t.Log(n)
}
