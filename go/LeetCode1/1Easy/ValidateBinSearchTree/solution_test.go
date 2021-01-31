package ValidateBinSearchTree

import "testing"

var cases = []struct {
	name   string
	inputs []int
	expect bool
}{
	{
		"test 1",
		[]int{2, 1, 3},
		true,
	},
	{
		"test 2",
		[]int{5, 1, 4, -1, -1, 3, 6},
		false,
	},
}

func generateTree(arr []int) *TreeNode {

}
func TestSolution(t *testing.T) {
	for _, c := range cases {
		r := generateTree(c.inputs)
		got := isValidBST(r)
		if got != c.expect {
			t.Fatalf("%s incorrect by %s", c.name, got)
		}
	}
}
