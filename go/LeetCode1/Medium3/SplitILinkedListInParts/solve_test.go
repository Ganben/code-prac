package splitilinkedlistinparts

import "testing"

func TestSolve(t *testing.T) {
	// head := []int{1, 2, 3}
	k := 5
	res := splitListToParts(head, k)
	// t.Log(res)
	if res[0].Val != 1 {
		t.FailNow()
	}
}
