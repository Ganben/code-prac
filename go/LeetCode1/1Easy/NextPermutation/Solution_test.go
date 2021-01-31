package Solution

import (
	"reflect"
	"testing"
)

var cases = []struct {
	name string
	inputs []int
	expect []int
}{
	{
		"test 1", []int{1, 2, 3}, []int{1, 3, 2},
	},
	{
		"test 2", []int{3, 2, 1}, []int{1, 2, 3},
	},
	{
		"test 3", []int{1, 1, 5}, []int{1, 5, 1},
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			nextPermutation(c.inputs)
			if !reflect.DeepEqual(c.inputs, c.expect) {
				t.Fatalf("expect %v got %v", c.expect, c.inputs)
			}
		})
	}

}