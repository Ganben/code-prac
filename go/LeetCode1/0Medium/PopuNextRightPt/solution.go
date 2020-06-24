package PopuNextRightPt

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	bfs([]*Node{root})
	return root
}

func bfs(arr []*Node) {
	if len(arr) == 0 {
		return
	}
	nextLevel := []*Node{}
	for i, n := range arr {
		if n != nil {
			nextLevel = append(nextLevel, n.Left)
			nextLevel = append(nextLevel, n.Right)
			if i < len(arr)-1 {
				n.Next = arr[i+1]
			} else {
				n.Next = nil
			}
		}
	}
	bfs(nextLevel)
}
