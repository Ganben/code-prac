package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1, root2 *TreeNode) bool {
	arr1 := []int{}
	dfs(root1, &arr1)
	arr2 := []int{}
	dfs(root2, &arr2)
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

func dfs(curr *TreeNode, arr *[]int) {
	if curr == nil {
		return
	}
	if curr.Left == nil && curr.Right == nil {
		*arr = append(*arr, curr.Val)
	} else {
		dfs(curr.Left, arr)
		dfs(curr.Right, arr)
	}
}
