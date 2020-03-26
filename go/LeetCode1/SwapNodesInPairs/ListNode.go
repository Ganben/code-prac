package Solution

import (
	// "fmt"
	// "math/rand"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func isEqual(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 == nil && l2 == nil {
		return false
	}
	if l1 != nil && l2 == nil {
		return false
	}
	return false
}