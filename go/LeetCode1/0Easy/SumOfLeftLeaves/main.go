package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *Treenode) int {
	// bfs
	l := []*TreeNode{}
	if root.Left != nil {
		l = append(l, root.Left)
	}
	if root.Right != nil {
		l = append(l, root.Right)
	}

	cur := 0
	if root.Left != nil {
		cur += root.Left.Val
	}
	for _, e := range l {
		cur += sumOfLeftLeaves(e)
	}
	return cur

}
