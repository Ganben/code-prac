package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	curr := root
	result := []int{}
	if curr == nil {
		return []int{}
	}
	result = append(result, curr.Val)
	ar := levelAdder(curr)
	ifnil, va := levelSearcher(ar)
	if ifnil {
		result = append(result, va)
	}
	for ifnil == true {
		// build next levels
		arrs := []*TreeNode{}
		for _, node := range ar {
			arrs = append(arrs, levelAdder(node)...)
		}
		ifnil, va := levelSearcher(arrs)
		if ifnil {
			result = append(result, va)
			ar = arrs
		} else {
			break
		}
	}
	return result
}

func levelAdder(parentNode *TreeNode) []*TreeNode {
	if parentNode.Left != nil && parentNode.Right != nil {
		ar := []*TreeNode{}
		ar = append(ar, parentNode.Right)
		ar = append(ar, parentNode.Left)
		return ar
	}
	return nil
}

func levelSearcher(ar []*TreeNode) (bool, int) {
	for _, t := range ar {
		if t != nil {
			return true, t.Val
		}
	}
	return false, 0
}
