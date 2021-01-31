package BinTreePostOrderTrav

func postorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	output := []int{}
	if root == nil {
		return output
	}
	stack = append(stack, root)
	next := false
	for !next {
		node := stack[len(stack)-1]
		// result incorrect
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		next = node.Left == nil && node.Right == nil
	}
	for i := len(stack) - 1; i >= 0; i-- {
		n := stack[i]
		output = append(output, n.Val)
	}
	return output

}
