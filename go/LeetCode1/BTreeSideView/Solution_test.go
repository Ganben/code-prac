package solution

import "testing"

var cases = []struct {
	name   string
	inputs []int
	expect []int
}{
	{
		"test 1",
		[]int{1, 2, 3, -1, 5, -1, 4},
		[]int{1, 3, 4},
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		l := 1
		root := &TreeNode{}
		pNodes := []*TreeNode{}
		childNodes := []*TreeNode{}
		for i := 0; i < len(c.inputs); i++ {
			pNodes = childNodes
			childNodes = []*TreeNode{}
			for j := 0; j < l; j++ {

				if i == 0 {
					// add root
					root = &TreeNode{
						Val:   c.inputs[i],
						Left:  nil,
						Right: nil,
					}
					childNodes = append(childNodes, root)
					l = 2
					continue
				}
				if c.inputs[i]
				nNode := &TreeNode{

				}
				childNodes = append(childNodes, nNode)

				l = l * 2
			}
			// assign parent-child
			if l > 0 {
				for _, el := range pNodes {
					
				}
			}

		}
		got := rightSideView()
	}
}
