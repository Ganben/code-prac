package ReorderList

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	p := head
	q := head
	r := head
	s := head
	for q.Next != nil {
		q = q.Next
		p = p.Next
		if q.Next != nil {
			q = q.Next
		}
	}
	q = p.Next
	p.Next = nil
	for q != nil {
		r = q.Next
		q.Next = p.Next
		p.Next = q
		q = r
	}
	q = p.Next
	p.Next = nil
	for q != nil {
		r = q.Next
		q.Next = s.Next
		s.Next = q
		s = q.Next
		q = r
	}
}
