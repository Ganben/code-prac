package solution

import "testing"

var cases = []struct {
	name   string
	input  [][]int
	expect int
}{
	{
		"test 1",
		[][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		},
		2,
	},
	{
		"test 2",
		[][]int{
			{0},
			{0},
		},
		1,
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		got := uniquePathsWithObstacles(c.input)
		if got != c.expect {
			t.Fatalf("%s incorrect %v", c.name, got)
		}

	}
}
