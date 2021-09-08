package constructrectangle

import "testing"

func TestSolve(t *testing.T) {
	res := constructRectangle(4)
	t.Log(res)
	if res[0]*res[1] != 4 {
		t.FailNow()
	}
}
