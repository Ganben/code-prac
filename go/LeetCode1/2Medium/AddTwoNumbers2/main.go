package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1, s2 := []int{}, []int{}
	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}
	carry := 0
	var ans *ListNode
	for len(s1) > 0 || len(s2) > 0 {
		carry += pop(&s1) + pop(&s2)
		ans, carry = &ListNode{Val: carry % 10, Next: ans}, carry/10
	}
	if carry != 0 {
		ans = &ListNode{Val: carry, Next: ans}
	}
	return ans
}

func pop(stack *[]int) int {
	ans := 0
	if len(*stack) > 0 {
		ans = (*stack)[len(*stack)-1]
		*stack = (*stack)[0 : len(*stack)-1]
	}
	return ans
}
