package LinkListCycleII

func getIntersect(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return slow
		}
	}
	return nil
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	intersect := getIntersect(head)
	if intersect == nil {
		return nil
	}
	pt1 := head
	pt2 := intersect
	for pt1 != pt2 {
		pt1 = pt1.Next
		pt2 = pt2.Next
	}
	return pt1
}
