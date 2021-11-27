package randomflipmatrix

import "testing"

func TestSolve(t *testing.T) {
	obj := Constructor(4, 5)
	param_1 := obj.Flip()
	if param_1 == nil {
		t.FailNow()
	}
	obj.Reset()
}
