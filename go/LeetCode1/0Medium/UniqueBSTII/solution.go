package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// var ans []*TreeNode

func generateTrees(n int) []*TreeNode {
	if n < 1 {
		return nil
	}

	return generateRec(1, n)
}

func generateRec(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{}
	}
	ans := []*TreeNode{}
	for i := start; i <= end+1; i++ {
		left_trees := generateRec(1, i-1)
		right_trees := generateRec(i+1, end+1)
		for _, l := range left_trees {
			for _, r := range right_trees {
				cur := &TreeNode{}
				cur.Val = i
				cur.Left = l
				cur.Right = r
				ans = append(ans, cur)
			}
		}
	}
	return ans
}
