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
	{
		"testcase 1",
		[][]int{{4, 5, 6, 7, 0, 1, 2}, {0}},
		4,
	},
	{
		"testcase 2",
		[][]int{{4, 5, 6, 7, 0, 1, 2}, {3}},
		-1,
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := search(c.inputs[0], c.inputs[1][0])
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expect %v got %v", c.expect, got)
			}
		})
	}
}