package Solution

import (
	"testing"
)

var cases = []struct {
	name string
	inputs []int
	expect int
}{
	{
		"test 1",
		[]int{2, 3, 1, 1, 4},
		2,
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		got := jump(c.inputs)
		if got != c.expect {
			t.Fatalf("incorrect %s", c.name)
		}
	}
}