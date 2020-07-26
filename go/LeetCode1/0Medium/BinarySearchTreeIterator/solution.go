package BinarySearchTreeIterator

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	nodesSorted []int
	index       int
}

func Constructor(root *TreeNode) BSTIterator {
	n := []int{}
	i := -1
	n = inorder(root)
	return BSTIterator{n, i}
}

func inorder(root *TreeNode) []int {
	res := []int{}
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		helper(root.Left)
		res = append(res, root.Val)
		helper(root.Right)
	}
	return res
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	if !this.HasNext() {
		return 0
	}
	this.index++
	i := this.index
	return this.nodesSorted[i]

}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.index+1 < len(this.nodesSorted)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
