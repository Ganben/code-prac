package solution

import "testing"

var cases = []struct {
	name   string
	inputs []int
	expect int
}{
	{
		"test 1",
		[]int{1, 8, 6, 2, 5, 4, 8, 3, 7},
		49,
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		got := maxArea(c.inputs)
		if got != c.expect {
			t.Fatalf("%s incorrect by %v", c.name, got)
		}
	}
}
