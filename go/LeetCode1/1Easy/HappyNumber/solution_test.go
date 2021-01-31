package solution

import "testing"

var cases = []struct {
	name string
	input int
	expect bool
}{
	{
		"test 1",
		19,
		true,
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		got := isHappy(c.input)
		if got != c.expect {
			t.Fatalf("%s incorrect by %v", c.name, got)
		}
	}
}