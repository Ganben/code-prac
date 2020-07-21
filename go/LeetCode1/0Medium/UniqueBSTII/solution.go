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
		return []*TreeNode{nil}
	}
	ans := []*TreeNode{}
	for i := start; i <= end; i++ {
		leftTrees := generateRec(1, i-1)
		rightTrees := generateRec(i+1, end)
		for _, l := range leftTrees {
			for _, r := range rightTrees {
				cur := &TreeNode{i, nil, nil}

				cur.Left = l
				cur.Right = r
				ans = append(ans, cur)
			}
		}
	}
	return ans
}
