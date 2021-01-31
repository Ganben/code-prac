package solution

import "testing"

func TestSolution(t *testing.T) {
	got := reversePairs([]int{7, 5, 6, 4})
	if got != 5 {
		t.Fatalf("incorrect by %v", got)
	}
}
