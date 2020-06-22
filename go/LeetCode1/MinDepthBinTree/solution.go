package solution

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	children := []*TreeNode{root.Left, root.Right}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	m := math.MaxInt32
	for _, node := range children {
		if node != nil {
			m = min(minDepth(node), m)
		}
	}
	return m + 1

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
