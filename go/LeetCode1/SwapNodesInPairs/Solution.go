package Solution

func swapPairs(head *ListNode) *ListNode {
	preHead := &ListNode{Val: 0, Next: nil}
	cur := preHead
	preHead.Next = head
	for cur.Next != nil && cur.Next.Next != nil {
		tmp := cur.Next.Next
		cur.Next.Next = tmp.Next
		tmp.Next = cur.Next
		cur.Next = tmp
		cur = cur.Next.Next
	}
	return preHead.Next
}

