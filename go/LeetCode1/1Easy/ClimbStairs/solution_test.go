package solution

import "testing"

var cases = []struct {
	name   string
	input  int
	expect int
}{
	{
		"test 1",
		2,
		2,
	},
	{
		"test 2",
		3,
		3,
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		got := climbStairs(c.input)
		if got != c.expect {
			t.Fatalf("%s incorrect by %v", c.name, got)
		}
	}
}
