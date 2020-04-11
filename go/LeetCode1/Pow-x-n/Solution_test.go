package Solution

import (
	"testing"
)

var cases = []struct {
	name   string
	input1 float64
	input2 int
	expect float64
}{
	{
		"case 1",
		2.0,
		10,
		1024.0,
	},
	{
		"case 2",
		2.1,
		3,
		9.2610,
	},
	{
		"case 3",
		2.0,
		-2,
		0.25,
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		got := myPow(c.input1, c.input2)
		if got-c.expect > 0.0000001*c.expect {
			t.Fatalf("%s failed with %v", c.name, got)
		}
	}
}
