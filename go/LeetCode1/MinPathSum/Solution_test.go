package solution

import "testing"

var cases = []struct {
	name   string
	inputs [][]int
	expect int
}{
	{
		"test 1",
		[][]int{
			{1, 3, 1},
			{1, 5, 1},
			{4, 2, 1},
		},
		7,
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		got := minPathSum(c.inputs)
		if got != c.expect {
			t.Fatalf("%s incorrect by %v", c.name, got)
		}
	}
}
