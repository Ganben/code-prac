package SortList

func sortList(head *ListNode) *ListNode {
	n := 0
	for i := head; i != nil; i = i.Next {
		n++
	}
	dummy := &ListNode{0, nil}
	dummy.Next = head

	for i := 1; i < n; i = 2 * i {
		begin := dummy
		for j := 0; j+i < n; j = j + 2*i {
			first := begin.Next
			second := first
			for k := 0; k < i; k++ {
				second = second.Next
			}
			f := 0
			s := 0
			for f < i && s < i && second != nil {
				if first.Val < second.Val {
					begin.Next = first
					begin = begin.Next
					first = first.Next
					f++
				} else {
					begin.Next = second
					begin = begin.Next
					second = second.Next
					s++
				}
			}
			for f < i {
				begin.Next = first
				begin = begin.Next
				first = first.Next
				f++
			}
			for s < i && second != nil {
				begin.Next = second
				begin = begin.Next
				second = second.Next
				s++
			}
			begin.Next = second
		}

	}
	return dummy.Next
}
