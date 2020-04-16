package solution

import (
	"reflect"
	"testing"
)

var cases = []struct {
	name    string
	inputs1 [][]int
	inputs2 []int
	expect  [][]int
}{
	{
		"test 1",
		[][]int{
			{1, 3},
			{6, 9},
		},
		[]int{2, 5},
		[][]int{
			{1, 5},
			{6, 9},
		},
	},
	{
		"test 2",
		[][]int{
			{2, 5},
			{6, 7},
			{8, 9},
		},
		[]int{0, 1},
		[][]int{
			{0, 1},
			{2, 5},
			{6, 7},
			{8, 9},
		},
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		got := insert(c.inputs1, c.inputs2)
		if !reflect.DeepEqual(got, c.expect) {
			t.Fatalf("%s incorrect by %v", c.name, got)
		}
	}
}
