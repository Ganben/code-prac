package mostfrequentsubtreesum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var hMap map[int]int

func findFrequentTreeSum(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	hMap = map[int]int{}
	treeSum(root, hMap)
	for k, v := range hMap {
		switch {
		case len(ret) == 0:
			ret = append(ret, k)
		case v > m[ret[0]]:
			ret = ret[:1]
			ret[0] = k
		case v == hMap[ret[0]]:
			ret = append(ret, k)
		}
	}
	return ret

}

func treeSum(root *TreeNode, m map[int]int) int {
	switch {
	case root.Left == nil && root.Right == nil:
		m[root.Val]++
		return root.Val
	default:
		sum := root.Val
		if root.Left != nil {
			sum += treeSum(root.Left, m)
		}
		if root.Right != nil {
			sum += treeSum(root.Right, m)

		}
		m[sum]++
		return sum
	}
}
