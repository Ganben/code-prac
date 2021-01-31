package RemoveLinkedListEle

func removeElements(head *ListNode, val int) *ListNode {
	cur1, cur2 := &ListNode{}, &ListNode{}
	if head == nil {
		return nil
	}
	cur1 = head
	// cur2 = head.Next

	// first no val
	for cur1.Val == val {
		if cur1.Next != nil {
			cur1 = cur1.Next
		} else {
			return nil
		}
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
