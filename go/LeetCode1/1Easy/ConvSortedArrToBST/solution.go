package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var numsArr []int

func sortedArrayToBST(nums []int) *TreeNode {
	numsArr = nums

	return helper(0, len(nums)-1)
}

func helper(left, right int) *TreeNode {
	if left > right {
		return nil
	}
	p := (left + right) / 2
	if (left+right)%2 == 1 {
		p += 1
	}
	root := &TreeNode{numsArr[p], nil, nil}
	root.Left = helper(left, p-1)
	root.Right = helper(p+1, right)
	return root
}
