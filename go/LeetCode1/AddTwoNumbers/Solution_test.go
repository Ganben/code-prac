package Solution

import (
	"testing"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		name string
		input1 *ListNode
		input2 *ListNode
		expect *ListNode
	}{
		{"1 test 1", &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}},
	&ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}},
	&ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8, Next: nil}}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := addTwoNumber(c.input1, c.input2)
			if !isEqual(got, c.expect) {
				t.Fatalf("expected %v, got %v", c.expect, got)
			}
		})
	}
}