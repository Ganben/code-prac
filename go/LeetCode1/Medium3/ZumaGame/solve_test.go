package zumagame

import "testing"

func TestSolve(t *testing.T) {
	b := "WRRBBW"
	h := "RB"
	res := findMinStep(b, h)
	if res != -1 {
		t.Log(res)
		t.FailNow()
	}

	b = "WWBBWBBWW"
	h = "BB"
	res = findMinStep(b, h)
	if res != -1 {
		t.Log(res)
		t.FailNow()
	}
}
