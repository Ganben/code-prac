package Solution

import (
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		got := isPalindrome(121121)
		expect := true
		if expect != got {
			t.Error("Got:", got, "expect:", expect)
		}
	})

	t.Run("test 2", func(t *testing.T) {
		got := isPalindrome2(-121)
		expect := false
		if got != expect {
			t.Fatalf("expect %v but got %v", expect, got)
		}
	})

	
}