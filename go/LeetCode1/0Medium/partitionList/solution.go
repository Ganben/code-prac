package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	//double pointer
	before := &ListNode{0, nil}
	b1 := before
	after := &ListNode{0, nil}
	a1 := after

	for head != nil {
		if head.Val < x {
			before.Next = head
			before = before.Next
		} else {
			after.Next = head
			after = after.Next
		}
		head = head.Next
	}
	after.Next = nil
	before.Next = a1.Next
	return b1.Next
}
