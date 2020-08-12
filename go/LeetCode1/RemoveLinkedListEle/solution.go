package RemoveLinkedListEle

func removeElements(head *ListNode, val int) *ListNode {
	cur1, cur2 := &ListNode{}, &ListNode{}

	cur1 = head
	// cur2 = head.Next

	// first no val
	for cur1.Val == val {
		cur1 = cur1.Next
	}
	origin := cur1
	cur2 = cur1.Next
	for cur2 != nil {
		if cur2.Val == val {
			cur1.Next = cur2.Next
			cur2 = cur2.Next
		} else {
			cur1 = cur2
			cur2 = cur2.Next
		}
	}

	return origin

}
