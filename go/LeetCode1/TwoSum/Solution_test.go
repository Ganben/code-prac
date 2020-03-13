package Solution

import (
	// "math/rand"
	"reflect"
	"testing"
)

func TestTowSum1(t *testing.T) {
	//
	cases := []struct {
		name string
		inputs [][]int
		expect []int
	}{
		{"1 test 1", [][]int{{2,7,11,15}, {9}}, []int{0,1}},
		{"2 test 2", [][]int{{3,2,4}, {6}}, []int{1,2}},
		{"3 test 3", [][]int{{2,7,11,15}, {9}}, []int{0,1}},
		{"4 test 4", [][]int{{7,6,7,3,2,1,4,9,10}, {17}}, []int{0,8}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := TowSum1(c.inputs[0], c.inputs[1][0])
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, got %v", c.expect, ret)
			}
		})
	}

}