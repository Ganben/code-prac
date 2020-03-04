package StackArray

import "testing"

func TestStack(t *testing.T) {
	// b := make([]int, 10)
	s := Stack{}
	s.Push(1)
	s.Push(99)
	if v, ok := s.Pop(); ok {
		if v != 99 {
			t.Fail()
		}
		if s.top != 1 {
			t.Fail()
		}
	}
}