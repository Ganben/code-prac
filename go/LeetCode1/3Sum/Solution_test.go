package Solution

import (
	"reflect"
	"testing"
)

var cases = []struct {
	name string
	inputs []int
	expect [][]int
}{
	{"test 1", []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
	{"test 2", []int{-2, 0, 0, 2, 2}, [][]int{{-2, 0, 2}}},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := threeSum(c.inputs)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect %v got %v", c.expect, ret)
			}
		})
	}

}