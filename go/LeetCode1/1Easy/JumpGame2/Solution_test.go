package Solution

import (
	"testing"
)

var cases = []struct {
	name   string
	inputs []int
	expect int
}{
	{
		"test 1",
		[]int{2, 3, 1, 1, 4},
		2,
	},
	{
		"test 2",
		[]int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3},
		2,
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		got := jump(c.inputs)
		if got != c.expect {
			t.Fatalf("incorrect %s by %v", c.name, got)
		}
	}
}
