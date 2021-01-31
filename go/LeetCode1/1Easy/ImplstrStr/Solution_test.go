package Solution

import "testing"

func TestSolution(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		got := strStr("hello", "ll")
		want := 2
		if got != want {
			t.Fatalf("expect %v got %v", want, got)
		}
	})

	t.Run("test 2", func(t *testing.T) {
		got := strStr("aaaaa", "ba")
		want := -1
		if got != want {
			t.Fatalf("expect %v got %v", want, got)
		}
	})
}