package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var x *TreeNode
var y *TreeNode
var pred *TreeNode

func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}
	x, y, pred = nil, nil, nil
	//&TreeNode{}, &TreeNode{}, &TreeNode{}
	findTwoSwapped(root)
	x.Val, y.Val = y.Val, x.Val
}

func findTwoSwapped(root *TreeNode) {
	if root == nil {
		return
	}

	findTwoSwapped(root.Left)
	if pred != nil && (root.Val < pred.Val) {
		y = root
		if x == nil {
			x = pred
		} else {
			return
		}
	}
	pred = root
	findTwoSwapped(root.Right)

}
