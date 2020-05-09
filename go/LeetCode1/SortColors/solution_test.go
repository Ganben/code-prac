package SortColors

import (
	"reflect"
	"testing"
)

var cases = []struct {
	name   string
	inputs []int
	expect []int
}{
	{
		"test 1",
		[]int{2, 0, 2, 1, 1, 0},
		[]int{0, 0, 1, 1, 2, 2},
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		sortColors(c.inputs)
		if !reflect.DeepEqual(c.inputs, c.expect) {
			t.Fatalf("%s incorrect by %v", c.name, c.inputs)
		}
	}
}
