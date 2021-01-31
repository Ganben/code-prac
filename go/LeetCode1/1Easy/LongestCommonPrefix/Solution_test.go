package Solution

import "testing"

var cases = []struct {
	name string
	inputs []string
	expect string
}{
	{"test 1", []string{"flower", "flow", "floor"}, "flo"},
	{"test 2", []string{"di", "echo", "car"}, ""},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		got := longestCommonPrefix(c.inputs)
		if got != c.expect {
			t.Fatalf("expected %v got %v", c.expect, got)
		}
	}
}