package Solution

import "testing"

func TestSolution(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		data := []int{1, 1, 2}
		got := removeDuplicates(data)
		want := 2
		if got != want {
			t.Fatalf("expecte %v got %v", want, got)
		}
	})
}
