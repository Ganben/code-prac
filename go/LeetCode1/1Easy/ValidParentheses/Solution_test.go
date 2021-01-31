package Solution

import "testing"

var cases = []struct {
	name string
	inputs string
	expect bool
}{
	{
		"test 1",
		"()",
		true,
	},
	{
		"test 2",
		"()[]{}",
		true,
	},
	{
		"test 3",
		"({[]})",
		true,
	},
	{
		"test 4",
		"({}}",
		false,
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		ret := isValid(c.inputs)
		if ret != c.expect {
			t.Fatalf("expect %v got %v in %s", c.expect, ret, c.name)
		}
	}
}