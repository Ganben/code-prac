package FlattenBinTreeToLinkedList

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	var arr []*TreeNode
	arr = []*TreeNode{}
	var traver func(*TreeNode)
	traver = func(node *TreeNode) {
		if node == nil {
			return
		}
		arr = append(arr, node)
		traver(node.Left)
		traver(node.Right)
	}
	traver(root)
	for i, node := range arr {
		node.Left = nil
		if i < len(arr)-1 {
			node.Right = arr[i+1]
		} else {
			node.Right = nil
		}
	}

}
