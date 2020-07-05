package CloneGraph

type Node struct {
	Val       int
	Neighbors []*Node
}

var visited map[*Node]*Node

func cloneGraph(node *Node) *Node {
	visited = map[*Node]*Node{}

	return helper(node)

}

func helper(node *Node) *Node {

	if node == nil {
		return node
	}
	if v, ok := visited[node]; ok {
		return v
	}
	cloneNode := &Node{node.Val, []*Node{}}
	visited[node] = cloneNode
	if len(node.Neighbors) > 0 {
		for _, n := range node.Neighbors {
			cloneNode.Neighbors = append(cloneNode.Neighbors, cloneGraph(n))
		}

	}
	return cloneNode
}
