package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	var helper func(*TreeNode) (bool, int)
	helper = func(node *TreeNode) (bool, int) {
		if node == nil {
			return true, -1
		}
		leftIsBalanced, leftHeight := helper(root.Left)
		if leftIsBalanced == false {
			return false, 0
		}
		rightIsBalanced, rightHeight := helper(root.Right)
		if rightIsBalanced == false {
			return false, 0
		}
		return abs(leftHeight-rightHeight) < 2, 1 + max(leftHeight, rightHeight)
	}
	res, _ := helper(root)
	return res
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
