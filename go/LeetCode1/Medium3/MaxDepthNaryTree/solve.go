package maxdepthnarytree

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	cnt := 0
	maxD := 0
	curr := root
	if len(curr.Children) > 0 {
		cnt += 1

		for _, v := range curr.Children {
			if cd := maxDepth(v); cd > maxD {
				maxD = maxDepth(v)
			}
		}
	}
	return cnt + maxD
}

func maxDepth2(root *Node) int {
	if root == nil {
		return 0
	}
	cnt := 0
	queue := []*Node{root}
	for len(queue) > 0 {
		q := queue
		queue = nil
		for _, node := range q {
			queue = append(queue, node.Children...)
		}
		cnt++
	}
	return cnt
}
