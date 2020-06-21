package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var intArr []int

func sortedListToBST(head *ListNode) *TreeNode {
	intArr = mapListToArr(head)
	return convert(0, len(intArr)-1)
}

func mapListToArr(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res

}

func convert(l, r int) *TreeNode {
	if l > r {
		return nil
	}
	mid := (l + r) / 2
	node := &TreeNode{intArr[mid], nil, nil}
	if l == r {
		return node
	}
	node.Left = convert(l, mid-1)
	node.Right = convert(mid+1, r)
	return node
}
