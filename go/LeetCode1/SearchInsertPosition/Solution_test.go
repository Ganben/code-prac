package Solution

import (
	"testing"
	"reflect"
)

var cases = []struct {
	name string
	inputs []int
	args int
	expect int
}{
	{
		"test 1",
		[]int{1, 3, 5, 6},
		5,
		2,
	},
	{
		"test 2",
		[]int{1, 3, 5, 6},
		2,
		1,
	},
	{
		"test 3",
		[]int{1, 3, 5, 6},
		7,
		4,
	},
	{
		"test 4",
		[]int{1, 3, 5, 6},
		0,
		0,
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := searchInsert(c.inputs, c.args)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expect %v got %v", c.expect, got)
			}
		})
	}
}