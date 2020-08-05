package HouseRobber3

// * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	f map[*TreeNode]int
	g map[*TreeNode]int
)

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	f, g = map[*TreeNode]int{}, map[*TreeNode]int{}
	dfs(root)
	return max(f[root], g[root])
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	dfs(root.Right)
	f[root] = root.Val + g[root.Left] + g[root.Right]
	g[root] = max(f[root.Left], g[root.Left]) + max(f[root.Right], g[root.Right])

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
