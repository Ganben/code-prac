package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans []bool

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// recurrent walker
	ans = []bool{}
	nodeEqual(p, q)
	for _, v := range ans {
		if v == false {
			return false
		}
	}

	return true
}

func nodeEqual(p *TreeNode, q *TreeNode) {
	if p != nil && q != nil {
		if p.Val != q.Val {
			ans = append(ans, false)
			return
		} else {
			// left
			if p.Left != nil && q.Left != nil {
				nodeEqual(p.Left, q.Left)
			} else {
				if !(p.Left == nil && q.Left == nil) {
					ans = append(ans, false)
				}

			}
			// right
			if p.Right != nil && q.Right != nil {
				nodeEqual(p.Right, q.Right)
			} else {
				if !(p.Right == nil && q.Right == nil) {
					ans = append(ans, false)
				}

			}
		}

	} else {
		if p == nil && q == nil {
			return
		}
		ans = append(ans, false)
		return
	}

}
