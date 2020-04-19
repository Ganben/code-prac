package solution

import "testing"

var cases = []struct {
	name   string
	inputo []int
	input2 int
	expect int
}{
	{
		"test 1",
		[]int{1, 2, 3, 4, 5},
		2,
		4,
	},
	{
		"test 2",
		[]int{1},
		1,
		1,
	},
	{
		"test 3",
		[]int{1, 2},
		2,
		1,
	},
}

func TestRotation(t *testing.T) {
	var head *ListNode
	for _, c := range cases {
		var curr *ListNode
		for i, e := range c.inputo {
			if i == 0 {
				head = &ListNode{e, nil}
				curr = head
			} else {
				n := &ListNode{e, nil}
				curr.Next = n
				curr = n
			}
		}
		got := rotateRight(head, c.input2)
		if got.Val != c.expect {
			t.Fatalf("%s case incorrect for %v", c.name, got.Val)
		}
	}
}
