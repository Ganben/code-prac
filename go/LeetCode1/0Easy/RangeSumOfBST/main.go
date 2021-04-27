package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	level := []*TreeNode{}
	if root == nil {
		return 0
	}
	ans := 0
	if root.Val >= low && root.Val <= high {
		ans += root.Val
	}
	level = append(level, root.Left)
	level = append(level, root.Right)
	for len(level) > 0 {
		newLevel := []*TreeNode{}
		for _, nd := range level {
			if nd.Val >= low && nd.Val <= high {
				ans += nd.Val
			}
			if nd.Left != nil {
				newLevel = append(newLevel, nd.Left)

			}
			if nd.Right != nil {
				newLevel = append(newLevel, nd.Right)
			}
		}
		level = newLevel
	}
	return ans
}
