package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sumCnt map[int]int

func pathSum(root *TreeNode, targetSum int) int {
	sumCnt = make(map[int]int)
	sumCnt[0] = 1
	return dfs(root, targetSum, 0)
}

func dfs(tn *TreeNode, sum, nowSum int) int {
	if tn == nil {
		return 0
	}
	//
	var result int
	nowSum = nowSum + tn.Val
	result = result + sumCnt[nowSum-sum]
	//
	sumCnt[nowSum] = sumCnt[nowSum] + 1
	//
	result = result + dfs(tn.Left, sum, nowSum)
	result = result + dfs(tn.Right, sum, nowSum)
	sumCnt[nowSum] = sumCnt[nowSum] - 1
	return result
}
