package main

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return root
	}
	nilHead := &Node{0, nil, root, nil}
	dfs(nilHead, root)
	nilHead.Next.Prev = nil
	return nilHead.Next
}

func dfs(prev, curr *Node) *Node {
	if curr == nil {
		return prev
	}
	curr.Prev = prev
	prev.Next = curr
	tmpNext := curr.Next
	tail := dfs(curr, curr.Child)
	curr.Child = nil
	return dfs(tail, tmpNext)
}
