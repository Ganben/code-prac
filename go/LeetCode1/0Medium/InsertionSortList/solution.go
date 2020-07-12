package InsertionSortList

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{-1}
	dummy.Next = head
	pre := dummy.Next

	for pre.Next != nil {
		val := pre.Next.Val
		nex := pre.Next.Next
		pi := dummy
		for ; pi != pre; pi = pi.Next {
			if pi.Next.Val > val {
				pj := pi.Next
				swapNode := pre.Next
				pi.Next = swapNode
				swapNode.Next = pj
				pre.Next = nex
				break
			}
		}
		if pi == pre {
			pre = pre.Next
		}
		return dummy.Next
	}
}
