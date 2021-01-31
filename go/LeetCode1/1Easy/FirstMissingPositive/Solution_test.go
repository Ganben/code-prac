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
		[]int{1, 2, 0},
		3,
	},
	{
		"test 2",
		[]int{3, 4, -1, 1},
		2,
	},
	{
		"test 3",
		[]int{7, 8, 9, 11, 12},
		1,
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		got := firstMissingPositive(c.inputs)
		if got != c.expect {
			t.Fatalf("expect %v got %v", c.expect, got)
		}
	}
}