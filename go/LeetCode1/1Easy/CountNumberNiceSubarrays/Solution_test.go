package solution

import "testing"

var cases = []struct {
	name   string
	inputs []int
	inputk int
	expect int
}{
	{
		"test 1",
		[]int{1, 1, 2, 1, 1},
		3,
		2,
	},
	{
		"test 2",
		[]int{2, 4, 6},
		1,
		0,
	},
	{
		"test 3",
		[]int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2},
		2,
		16,
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		got := numberOfSubarrays(c.inputs, c.inputk)
		if got != c.expect {
			t.Fatalf("%s incorrect %v\n", c.name, got)
		}
	}
}
