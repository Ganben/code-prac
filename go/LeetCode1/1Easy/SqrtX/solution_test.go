package solution

import "testing"

var cases = []struct {
	name   string
	inputs int
	expect int
}{
	{
		"test 1",
		4,
		2,
	},
	{
		"test 2",
		8,
		2,
	},
	{
		"test 3",
		3,
		1,
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		got := mySqrt(c.inputs)
		if got != c.expect {
			t.Fatalf("%s case failed by %v", c.name, got)
		}
	}
}
