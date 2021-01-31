package Solution

import (
	"reflect"
	"testing"
)

var cases = []struct {
	name   string
	inputs []int
	expect [][]int
}{
	{
		"case 1",
		[]int{1, 1, 2},
		[][]int{
			{1, 1, 2},
			{1, 2, 1},
			{2, 1, 1},
		},
	},
	{
		"case 2",
		[]int{1},
		[][]int{
			{1},
		},
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		got := permuteUnique(c.inputs)
		if !reflect.DeepEqual(got, c.expect) {
			t.Fatalf("%s incorrect %v for %v", c.name, got, numset)
		}
	}
}
