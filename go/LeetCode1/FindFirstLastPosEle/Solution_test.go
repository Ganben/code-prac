package Solution

import (
	"reflect"
	"testing"
)

var cases = []struct {
	name string
	inputs [][]int
	expect []int
}{
	{
		"testcase 1",
		[][]int{{5, 7, 7, 8, 8, 10}, {8}},
		[]int{3, 4},
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := searchRange(c.inputs[0], c.inputs[1][0])
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expect %v got %v", c.expect, got)
			}
		})
	}
}