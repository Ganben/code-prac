package solution

var inorderArr []int
var postorderArr []int
var idxMap map[int]int

func buildTree(inorder []int, postorder []int) *TreeNode {
	inorderArr = inorder
	postorderArr = postorder
	idxMap = map[int]int{}
	for i, v := range inorder {
		idxMap[v] = i
	}
	return traver(0, len(inorder)-1)
}

func traver(inLeft, inRight int) *TreeNode {
	if inLeft > inRight {
		return nil
	}
	val := postorderArr[len(postorderArr)-1]
	postorderArr = postorderArr[:len(postorderArr)-1]
	root := &TreeNode{val, nil, nil}
	index := idxMap[val]

	//
	root.Right = traver(index+1, inRight)
	root.Left = traver(inLeft, index-1)
	return root
}
