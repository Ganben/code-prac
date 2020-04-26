package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil {
		return nil
	}
	n := len(lists)
	if n == 0 {
		return nil
	}
	return merge(lists, 0, n-1)
}

func merge(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	mid := int(left + (right-left)/2)
	l1 := merge(lists, left, mid)
	l2 := merge(lists, mid+1, right)
	return mergeTwoLists(l1, l2)
}

func mergeTwoLists(left, right *ListNode) *ListNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	if left.Val < right.Val {
		left.Next = mergeTwoLists(left.Next, right)
		return left
	}
	right.Next = mergeTwoLists(left, right.Next)
	return right
}
