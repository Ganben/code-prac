package Solution

import (
	"reflect"
	"testing"
)

var cases = []struct {
	name string
	inputs *ListNode
	expect *ListNode
}{
	{"1 test 1",
	&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}},
	&ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}},
},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := swapPairs(c.inputs)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("incorrect")
			}
		})
	}
}