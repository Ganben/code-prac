package solution

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Val: -1, Next: head}
	a := dummy
	b := head
	for b != nil && b.Next != nil {
		if a.Next.Val != b.Next.Val {
			a = a.Next
			b = b.Next
		} else {
			for b != nil && b.Next != nil && a.Next.Val == b.Next.Val {
				b = b.Next
			}
			a.Next = b.Next
			b = b.Next
		}
	}
	return dummy.Next
}
