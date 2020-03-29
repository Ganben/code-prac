package Solution

import (
	"reflect"
	"testing"
)

var cases = []struct {
	name string
	inputs []int
	expect int
}{
	{"1 test 1", []int{10, 3}, 3},
	{"2 test 2", []int{7, -3}, -2},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := divide(c.inputs[0], c.inputs[1])
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected %v got %v", c.expect, ret)
			}
		})
	}
}