package ValidateBinSearchTree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return search(root, math.MinInt64, math.MaxInt64)
}

func search(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return search(root.Left, lower, root.Val) && search(root.Right, root.Val, upper)
}
