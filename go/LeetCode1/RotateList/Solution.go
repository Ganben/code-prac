package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	var newHead *ListNode
	var tail *ListNode
	// var newTail *ListNode
	var curr *ListNode
	if head == nil {
		return nil
	}
	l := 1
	// length
	for tail = head; tail.Next != nil; tail = tail.Next {
		l++
	}
	// fmt.Printf("lenth count %v\n", l)
	if k >= l {
		k = k % l
	}
	if l == 1 {
		return head
	}
	if k == 0 {
		return head
	}
	var last *ListNode
	for i := 1; i < l+1; i++ {
		if i == 1 {
			last = nil
			curr = head
		} else {
			last = curr
			curr = curr.Next
		}

		if i == l-k+1 {
			// fmt.Printf("new head at %v\n", curr.Val)
			newHead = curr
			if last != nil {
				last.Next = nil
			}
		}
	}
	tail.Next = head
	// newTail.Next = nil
	return newHead
}
