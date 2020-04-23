package solution

import "testing"

var cases = []struct {
	name   string
	input  int
	expect int
}{
	{
		"test 1",
		5,
		2,
	},
	{
		"test 2",
		10,
		4,
	},
}

func TestWaysToChange(t *testing.T) {
	for _, c := range cases {
		got := waysToChange(c.input)
		if got != c.expect {
			t.Fatalf("%s incorrect by %v", c.name, got)
		}
	}
}
