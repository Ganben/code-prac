package Solution

import (
	"reflect"
	"testing"
)

var cases = []struct {
	name string
	inputs [][]int
	expect int
}{
	{"test 1", [][]int{{-1, 2, 1, -4}, {1}}, 2},
}

func TestSolution(t *testing.T) {
	//
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := threeSumClosest(c.inputs[0], c.inputs[1][0])
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected %v got %v", c.expect, ret)
			}
		})
	}
}