package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var results [][]int

func zigzagLevelOrder(root *TreeNode) [][]int {
	results = [][]int{}
	if root == nil {
		return results
	}
	dfs(root, 0)
	return results
}

func dfs(node *TreeNode, level int) {
	if level >= len(results) {
		results = append(results, []int{node.Val})
	} else {
		if level%2 == 0 {
			results[level] = append(results[level], node.Val)
		} else {
			results[level] = append(results[level], 0)
			results[level] = append(results[level], node.Val)
		}
	}

	if node.Left != nil {
		dfs(node.Left, level+1)
	}
	if node.Right != nil {
		dfs(node.Right, level+1)
	}

}
