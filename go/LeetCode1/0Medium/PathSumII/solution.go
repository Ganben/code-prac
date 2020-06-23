package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {

	if root == nil {
		return [][]int{}
	}
	ans := [][]int{}

	var helper func(*TreeNode, []int)
	helper = func(node *TreeNode, cur []int) {
		if node == nil {
			return
		}
		cur = append(cur, node.Val)
		if node.Left == nil && node.Right == nil {
			if sum == checksum(cur) {
				tmp := make([]int, len(cur))
				copy(tmp, cur)
				ans = append(ans, tmp)
				return
			}
			return
		}
		helper(node.Left, cur)
		helper(node.Right, cur)
	}
	helper(root, []int{})
	return ans
}

func checksum(arr []int) int {
	ans := 0
	for _, v := range arr {
		ans += v
	}
	return ans
}
