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
	{"test 1", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := maxArea(c.inputs)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect %v got %v", c.expect, ret)
			}
		})
	}
}