package solution

func maxDepth(root *TreeNode) int {
	stack := []*TreeNode{}
	depstack := []int{}
	if root == nil {
		return 0
	}
	stack = append(stack, root)
	depstack = append(depstack, 1)
	depth := 0
	curDepth := 0
	curNode := &TreeNode{}

	for len(stack) > 0 {
		curDepth = depstack[len(depstack)-1]
		depstack = depstack[:len(depstack)-1]
		curNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if curNode != nil {
			depth = max(depth, curDepth)
			stack = append(stack, curNode.Left)
			stack = append(stack, curNode.Right)
			depstack = append(depstack, curDepth+1)
			depstack = append(depstack, curDepth+1)
		}
	}
	return depth
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
