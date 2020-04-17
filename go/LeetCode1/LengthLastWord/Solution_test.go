package solution

import "testing"

var cases = []struct {
	name   string
	inputs string
	expect int
}{
	{
		"test 1",
		"a ",
		1,
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		got := lengthOfLastWord(c.inputs)
		if got != c.expect {
			t.Fatalf("%s case fail by %v", c.name, got)
		}
	}
}
