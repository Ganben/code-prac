package perfectnumber

import "testing"

func TestSolve(t *testing.T) {
	if f := checkPerfectNumber(28); !f {
		t.FailNow()
	}
}
