package nextgreaterelement2

import "testing"

func TestSolve(t *testing.T) {
	arr := []int{1, 2, 1}
	res := nextGreaterElements(arr)
	expected := []int{2, -1, 2}
	for i, v := range res {
		if v != expected[i] {
			t.FailNow()
		}
	}
}
