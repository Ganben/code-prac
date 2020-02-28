package ReverseArray

import (
	// "math/rand"
	"testing"
	// "time"
	// "sort"
)

func TestReverseArray(t *testing.T) {
	a := make([]int, 10)
	b := make([]int, 10)
	for i := range a {
		a[i] = i
	}
	copy(b, a)
	ReverseArray(a)
	for i := range b {
		if a[i] != b[9-i] {
			t.Fail()
		}
	}
}