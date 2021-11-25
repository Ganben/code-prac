package searchbst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	var ans *TreeNode
	ans = nil
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if root.Left != nil {
		ans = searchBST(root.Left, val)
		if ans != nil {
			return ans
		}
	}
	if root.Right != nil {
		ans = searchBST(root.Right, val)
		if ans != nil {
			return ans
		}
	}
	return ans

}
