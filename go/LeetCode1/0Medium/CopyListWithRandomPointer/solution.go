package CopyListWithRandomPointer

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	ptr := head
	for ptr != nil {
		newNode := &Node{ptr.Val, nil, nil}
		newNode.Next = ptr.Next
		ptr.Next = newNode
		ptr = newNode.Next
	}
	//
	ptr = head
	for ptr != nil {
		if ptr.Random != nil {
			ptr.Next.Random = ptr.Random.Next
		} else {
			ptr.Next.Random = nil
		}
		ptr = ptr.Next.Next
	}
	ptrOld := head
	ptrNew := head.Next
	headOld := head.Next
	for ptrOld != nil {
		ptrOld.Next = ptrOld.Next.Next
		if ptrNew.Next != nil {
			ptrNew.Next = ptrNew.Next.Next
		} else {
			ptrNew.Next = nil
		}
		ptrOld = ptrOld.Next
		ptrNew = ptrNew.Next
	}
	return headOld
}
