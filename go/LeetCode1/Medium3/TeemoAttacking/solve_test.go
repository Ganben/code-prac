package teemoattacking

import "testing"

func TestSolve(t *testing.T) {
	res := findPoisonedDuration([]int{1, 2}, 2)
	if res != 3 {
		t.Log(res)
		t.FailNow()
	}

}
