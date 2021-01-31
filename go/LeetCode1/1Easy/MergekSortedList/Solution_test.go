package Solution

import (
	"testing"
)

var testListNode = []*ListNode{
	UnmarshalListBySlice([]int{1, 4, 5}),
	UnmarshalListBySlice([]int{1, 3, 4}),
	UnmarshalListBySlice([]int{2, 6}),
}

func TestSolution(t *testing.T) {
	got := mergeKLists(testListNode)
	if got.Val != 1 || got.Next.Val != 1 {
		t.Fatalf("expect head wrong")
	}
}