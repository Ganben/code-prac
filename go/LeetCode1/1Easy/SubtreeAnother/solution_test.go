package SubtreeAnother

import "testing"

var cases = []struct {
	name   string
	input1 []int
	input2 []int
	expect bool
}{
	{
		"test 1",
		[]int{3, 4, 5, 1, 2, 0, 0},
		[]int{4, 1, 2},
		true,
	},
}

func generateTree(arr []int) *TreeNode {

}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		r1 := generateTree(c.input1)
		r2 := generateTree(c.input2)
		got := isSubtree(r1, r2)
		if got != c.expect {
			t.Fatalf("%s incorrect by %v", c.name, got)
		}
	}
}
