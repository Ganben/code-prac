package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	//find left
	//record path(last two)

	path := []*TreeNode{}
	var inorder func(*TreeNode)
	inorder = func(tn *TreeNode) {
		if tn != nil {
			inorder(tn.Left)
			path = append(path, tn)
			inorder(tn.Right)
		}
		return
	}
	inorder(root)
	//
	for i, tn := range path {
		if i < len(path)-1 {
			tn.Right = path[i+1]
			tn.Left = nil
		} else {
			tn.Left = nil
			tn.Right = nil
		}
	}
	return path[0]
}

func increasingBST(root *TreeNode) *TreeNode {
	vals := []int{}
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node != nil {
			inorder(node.Left)
			vals = append(vals, node.Val)
			inorder(node.Right)
		}
	}
	inorder(root)

	dummyNode := &TreeNode{}
	curNode := dummyNode
	for _, val := range vals {
		curNode.Right = &TreeNode{Val: val}
		curNode = curNode.Right
	}
	return dummyNode.Right
}
