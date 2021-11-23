package buddystring

import "testing"

func TestSolve(t *testing.T) {
	if f := buddyStrings("ab", "ba"); !f {
		t.FailNow()
	}
}
