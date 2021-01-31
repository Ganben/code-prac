package Solution

import (
	"reflect"
	"testing"
)

var cases = []struct {
	name string
	inputs [][]int
	expect float64
}{
	{"1 test 1", [][]int{{1,3}, {2}}, 2.0},
	{"2 test 2", [][]int{{1,2}, {3,4}}, 2.5},
}

func TestSolution(t *testing.T) {
	//
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := findMedianSortedArrays(c.inputs[0], c.inputs[1])
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected %v got %v", c.expect, ret)
			}
		})
	}
}