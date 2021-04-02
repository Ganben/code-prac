package main

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := [][]int{}
	currLayer := []*Node{root}
	for len(currLayer) > 0 {
		curr := []*Node{}
		ansTemp := []int{}
		for _, node := range currLayer {
			curr = append(curr, node.Children...)
			ansTemp = append(ansTemp, node.Val)
		}
		//
		ans = append(ans, ansTemp)
		currLayer = curr
	}
	return ans
}
