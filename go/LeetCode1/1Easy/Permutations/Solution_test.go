package Solution

import (
	"testing"
)

var cases = []struct {
	name string
	inputs []int
	expect [][]int
}{
	{
		"case 1",
		[]int{1, 2, 3},
		[][]int{
			{1,2,3},
			{1,3,2},
			{2,1,3},
			{2,3,1},
			{3,1,2},
			{3,2,1},
		},
	},
}
func TestSolution(t *testing.T) {
	for _, c := range cases {
		got := permute(c.inputs)
		if len(c.expect) != len(got) {
			t.Fatalf("case %s fail", c.name)
		}
	}
}