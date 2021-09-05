package predictthewinner

import "testing"

func TestSolve(t *testing.T) {
	if false != PredictTheWinner([]int{1, 5, 2}) {
		t.Log("failed")
		t.FailNow()
	}
	t.Log("pass")
}
