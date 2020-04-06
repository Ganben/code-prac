package Solution

import (
	"testing"
	// "reflect"
)

var cases = []struct {
	name string
	inputs string
	inputp string
	expect bool
}{
	{
		"case 1",
		"aa",
		"a",
		false,
	},
	{
		"case 2",
		"aa",
		"*",
		true,
	},
	{
		"case 3",
		"cb",
		"?a",
		false,
	},
	{
		"case 4",
		"adceb",
		"*a*b",
		true,
	},
	{
		"case 5",
		"acdcb",
		"a*c?b",
		false,
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		got := isMatch(c.inputs, c.inputp)
		if got != c.expect {
			t.Fatalf("%s incorrect", c.name)
		}
	}
}