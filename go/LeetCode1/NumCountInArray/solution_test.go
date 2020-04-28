package solution

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
		"t1",
		[]int{4, 1, 4, 6},
		[]int{1, 6},
	},
	{
		"t2",
		[]int{1, 2, 10, 4, 1, 4, 3, 3},
		[]int{2, 10},
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		g := singleNumbers(c.inputs)
		if !reflect.DeepEqual(g, c.expect) {
			t.Fatalf("%s failed by %v", c.name, g)
		}
	}
}
