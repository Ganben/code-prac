package SumRootOfLeafs

func sumNumbers(root *TreeNode) int {
	return traver(root, 0)
}

func traver(root *TreeNode, i int) int {
	if root == nil {
		return 0
	}
	temp := i*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return temp
	}
	return traver(root.Left, temp) + traver(root.Right, temp)
}
