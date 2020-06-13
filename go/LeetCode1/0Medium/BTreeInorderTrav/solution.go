package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result []int

func inorderTraversal(root *TreeNode) []int {
	result = []int{}
	traver(root)
	return result
}

func traver(root *TreeNode) {
	if root != nil {
		if root.Left != nil {
			traver(root.Left)
		}
		result = append(result, root.Val)
		if root.Right != nil {
			traver(root.Right)
		}
	}
}
