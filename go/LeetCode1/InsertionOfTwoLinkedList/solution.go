package InsertionOfTwoLinkedList

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1 := headA
	p2 := headB
	end1 := &ListNode{0, nil}
	end2 := &ListNode{0, nil}
	f1, f2 := 0, 0
	flag := true
	if p1 == nil || p2 == nil {
		return nil
	}
	for flag {
		//
		if p1 == p2 {
			return p1
		}
		//
		if p1.Next != nil {
			p1 = p1.Next
		} else {
			f1++
			end1 = p1
			p1 = headB
		}
		if p2.Next != nil {
			p2 = p2.Next
		} else {
			f2++
			end2 = p2
			p2 = headA
		}

		//
		if f1 == 2 && f2 == 2 {
			flag = false
			if end1 != end2 {
				return nil
			}
		}

	}

}
