package reversepairs

import "testing"

func TestSolve(t *testing.T) {
	arr := []int{1, 3, 2, 3, 1}
	res := reversePairs(arr)
	if res != 2 {
		t.Log(res)
		t.FailNow()
	}
}
