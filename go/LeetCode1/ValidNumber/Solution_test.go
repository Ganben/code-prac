package solution

import "testing"

var cases = []struct {
	inputs string
	expect bool
}{
	{
		" 0.1",
		true,
	},
	{
		"2e10",
		true,
	},
	{
		" 6e-1",
		true,
	},
	{
		" 99e2.5",
		false,
	},
	{
		"53.5e93",
		true,
	},
	{
		" ",
		false,
	},
	{
		"1 ",
		true,
	},
	{
		"1  ",
		true,
	},
	{
		".1",
		true,
	},
	{
		"1.",
		true,
	},
	{
		" -.",
		false,
	},
}

func TestIsNumberic(t *testing.T) {
	for _, c := range cases {
		got := isNumber(c.inputs)
		if got != c.expect {
			t.Fatalf("%s incorrect %v", c.inputs, got)
		}
	}
}
