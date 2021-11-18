package bintreetilt

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	ans := 0
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		sumLeft := dfs(node.Left)
		sumRight := dfs(node.Right)
		ans += abs(sumLeft - sumRight)
		return sumLeft + sumRight + node.Val
	}
	dfs(root)
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
