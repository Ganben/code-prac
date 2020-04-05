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
		[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
		6,
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		got := trap(c.inputs)
		if got != c.expect {
			t.Fatalf("incorrect got %v", got)
		}
	}
}