package Solution

import (
	"reflect"
	"testing"
)

var cases = []struct {
	name   string
	inputs [][]int
	expect []int
}{
	{
		"test 1",
		[][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
	},
	{
		"test 2",
		[][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		},
		[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
	},
	{
		"test 3",
		[][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 15, 16},
		},
		[]int{1, 2},
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := spiralOrder(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("%s case fail by %v", c.name, got)
			}
		})
	}
}
