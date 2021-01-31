package DeleteNodeInLinkedList

func deleteNode(node *ListNode) {
	if node != nil && node.Next != nil {
		node.Val = node.Next.Val
		tmp := node.Next
		node.Next = tmp.Next

	}
}
