package Solution

import (
	"reflect"
	"testing"
)

var cases = []struct {
	name   string
	inputs [][]int
	expect [][]int
}{
	{
		"test 1",
		[][]int{
			{0, 0, 0},
			{0, 1, 0},
			{1, 1, 1},
		},
		[][]int{
			{0, 0, 0},
			{0, 1, 0},
			{1, 2, 1},
		},
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		got := updateMatrix(c.inputs)
		if !reflect.DeepEqual(got, c.expect) {
			t.Fatalf("%s failed by %v", c.name, got)
		}
	}
}
