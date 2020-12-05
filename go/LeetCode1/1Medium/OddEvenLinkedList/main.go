package main

// package OddEvenLinkedList

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p1 := head

	// p1Tail := head
	p2 := head.Next
	if p2 == nil {
		return head
	}

	p2Init := p2
	// odd := 1
	// cur := head
	for p2.Next != nil {
		np1 := p2.Next
		np2 := p2
		if np1.Next != nil {
			np2 = np1.Next
		} else {
			p1.Next = np1
			p1 = np1
			break
		}
		p1.Next = np1
		p2.Next = np2
		if np1 == nil {
			break
		}
		if np2 == nil {
			break
		}
		p1 = np1
		p2 = np2
	}

	if p2Init != nil {
		p1.Next = p2Init

	}
	if p2.Next != nil {
		p2.Next = nil
	}
	return head

}

func main() {
	head := &ListNode{Val: 0, Next: nil}
	cur := head
	for i := 0; i < 5; i++ {

		cur.Val = i
		cur.Next = &ListNode{Val: i + 1, Next: nil}
		cur = cur.Next
	}
	cur = head
	for cur.Next != nil {
		fmt.Printf("%v->>", cur.Val)
		cur = cur.Next
	}
	head = oddEvenList(head)
	cur = head
	for cur.Next != nil {
		fmt.Printf("%v->", cur.Val)
		cur = cur.Next
	}
}
