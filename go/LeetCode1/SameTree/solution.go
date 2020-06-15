package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans bool

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// recurrent walker
	ans = true
	nodeEqual(p, q)
	return ans
}

func nodeEqual(p *TreeNode, q *TreeNode) {
	if p != nil && q != nil {
		if p.Val != q.Val {
			ans = false
			return
		} else {
			// left
			if p.Left != nil && q.Left != nil {
				nodeEqual(p.Left, q.Left)
			} else {
				if p.Left == nil && q.Left == nil {
					return
				}
				ans = false
				return
			}
			// right
			if p.Right != nil && q.Right != nil {
				nodeEqual(p.Right, q.Right)
			} else {
				if p.Right == nil && q.Right == nil {
					return
				}
				ans = false
				return
			}
		}

	} else {
		if p == nil && q == nil {
			return
		}
		ans = false
		return
	}

}
