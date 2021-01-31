package InvertBinTree

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	bfs(root)
	return root
}

func bfs(root *TreeNode) {
	if root != nil {
		l := root.Left
		r := root.Right
		root.Left = r
		root.Right = l
		bfs(l)
		bfs(r)
	}
}
