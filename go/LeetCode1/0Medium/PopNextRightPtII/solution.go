package solution

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	leftmost := root
	for leftmost != nil {
		prev, curr := &Node{}, leftmost
		leftmost = nil
		for curr != nil {
			prev, leftmost = processChild(curr.Left, prev, leftmost)
			prev, leftmost = processChild(curr.Right, prev, leftmost)
			curr = curr.Next
		}
	}
	return root
}

func processChild(childNode, prev, leftmost *Node) (*Node, *Node) {
	if childNode != nil {
		if prev != nil {
			prev.Next = childNode
		} else {
			leftmost = childNode
		}
		prev = childNode
	}
	return prev, leftmost
}
