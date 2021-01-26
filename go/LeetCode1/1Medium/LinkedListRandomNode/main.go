package main

import (
	"fmt"
	"math/rand"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	head   *ListNode
	length int
}

func Constructor(head *ListNode) Solution {
	theHead := head
	theLength := 0
	for head != nil {
		head = head.Next
		theLength++
	}
	return Solution{theHead, theLength}
}

func (this *Solution) GetRandom() int {
	p := rand.Intn(this.length)
	cur := this.head
	for i := 0; i < p; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func main() {
	p := &ListNode{1, nil}
	obj := Constructor(p)
	param1 := obj.GetRandom()
	fmt.Printf("%v", param1)
}
