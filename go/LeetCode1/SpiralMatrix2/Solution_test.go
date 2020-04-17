package solution

import (
	"reflect"
	"testing"
)

var cases = []struct {
	name   string
	inputs int
	expect [][]int
}{
	{
		"test 1",
		3,
		[][]int{
			{1, 2, 3},
			{8, 9, 4},
			{7, 6, 5},
		},
	},
	{
		"t 2",
		2,
		[][]int{
			{1, 2},
			{4, 3},
		},
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		got := generateMatrix(c.inputs)
		if !reflect.DeepEqual(got, c.expect) {
			t.Fatalf("%s wrong by %v", c.name, got)
		}
	}
}
