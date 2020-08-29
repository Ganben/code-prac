package KthSmallestInBST

func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}

	curr := root
	for true {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return curr.Val
		}
		curr = curr.Right
	}
	return 0
}
