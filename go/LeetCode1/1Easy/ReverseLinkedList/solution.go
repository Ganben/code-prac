package ReverseLinkedList

func reverseList(head *ListNode) *ListNode {
	// tail := head
	// tail := head
	if head == nil {
		return head
	}
	cur1 := head
	if cur1.Next == nil {
		return cur1
	}
	cur2 := cur1.Next
	if cur2.Next == nil {
		cur2.Next = cur1
		cur1.Next = nil
		return cur2
	}
	cur3 := cur2.Next

	head.Next = nil
	for cur2 != nil {
		cur2.Next = cur1
		cur1 = cur2
		cur2 = cur3
		if cur3 != nil {
			cur3 = cur3.Next
		}
	}
	return cur1
}
