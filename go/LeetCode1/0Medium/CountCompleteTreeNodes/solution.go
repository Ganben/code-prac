package CountCompleteTreeNodes

func countNodes(root *TreeNode) int {
	cur := []*TreeNode{root}
	count := 1
	if root == nil {
		return 0
	}
	for len(cur) > 0 {
		nextCur := []*TreeNode{}
		for _, e := range cur {
			if e.Left != nil {
				count++
				nextCur = append(nextCur, e.Left)
			}
			if e.Right != nil {
				count++
				nextCur = append(nextCur, e.Right)
			}

		}
		cur = nextCur
	}
	return count
}
