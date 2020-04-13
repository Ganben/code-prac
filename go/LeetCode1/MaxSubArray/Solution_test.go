package maxsubarray

import "testing"

var cases = []struct {
	name   string
	inputs []int
	expect int
}{
	{
		"case 1",
		[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		6,
	},
	{
		"case 2",
		[]int{-2, 1},
		1,
	},
	{
		"case 3",
		[]int{1, 2},
		3,
	},
	{
		"case 4",
		[]int{3, -3, 2, -3},
		3,
	},
	{
		"case 5",
		[]int{3, 2, -3, -1, 1, -3, 1, -1},
		5,
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := maxSubArray(c.inputs)
			if got != c.expect {
				t.Fatalf("%s case incorrect wity %v", c.name, got)
			}
		})
	}
}
