package solution

var target int
var ans bool

func hasPathSum(root *TreeNode, sum int) bool {
	ans = false
	target = sum
	traver(root, []int{})
	return ans
}

func traver(root *TreeNode, cur []int) {
	if root == nil {
		return
	}
	cur = append(cur, root.Val)
	if root.Right == nil && root.Left == nil {
		if target == arrsum(cur) {
			ans = true
			return
		} else {
			return
		}
	}
	arr := []*TreeNode{root.Left, root.Right}
	for _, node := range arr {
		if node != nil {
			traver(node, cur)
		}
	}

}

func arrsum(arr []int) int {
	ans := 0
	for _, v := range arr {
		ans += v
	}
	return ans
}
