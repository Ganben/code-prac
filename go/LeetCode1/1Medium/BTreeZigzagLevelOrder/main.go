type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans [][]int

func zigzagLevelOrder(root *TreeNode) [][]int {
	ans = [][]int{}
	firstRow := []*TreeNode{root}
	bfsZigzag(firstRow)
	return ans
}

func bfsZigzag(nodes []*TreeNode) {
	next := []*TreeNode{}
	thisLine := []int{}
	conFlag := false
	if len(ans)%2 == 0 {
		for i := 0; i < len(nodes); i++ {
			if nodes[i] != nil {
				conFlag = true
				thisLine = append(thisLine, nodes[i].Val)

			}
		}
	} else {
		for i := len(nodes) - 1; i >= 0; i-- {
			if nodes[i] != nil {
				conFlag = true
				thisLine = append(thisLine, nodes[i].Val)
				// next = append(next, nodes[i].Right)
				// next = append(next, nodes[i].Left)
			}
		}
	}
	for i := 0; i < len(nodes); i++ {
		if nodes[i] != nil {
			next = append(next, nodes[i].Left)
			next = append(next, nodes[i].Right)
		}
	}

	if conFlag {
		ans = append(ans, thisLine)
		bfsZigzag(next)
	}

}

func main() {

}