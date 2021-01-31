package solution

import (
	"reflect"
	"testing"
)

var cases = []struct {
	name   string
	inputs [][]int
	expect [][]int
}{
	{
		"test 1",
		[][]int{
			{1, 3},
			{2, 6},
			{8, 10},
			{15, 18},
		},
		[][]int{
			{1, 6},
			{8, 10},
			{15, 18},
		},
	},
	{
		" test 2",
		[][]int{
			{1, 4},
			{0, 4},
		},
		[][]int{
			{0, 4},
		},
	},
}

func TestMerge(t *testing.T) {
	for _, c := range cases {
		got := merge(c.inputs)
		if !reflect.DeepEqual(got, c.expect) {
			t.Fatalf("case %s incorrect by %v", c.name, got)
		}
	}
}

func TestMergesort(t *testing.T) {
	got := mergeSort([][]int{{1, 4}, {0, 4}})
	if !reflect.DeepEqual(got, [][]int{{0, 4}, {1, 4}}) {
		t.Fatalf("merge sort fail by %v", got)
	}
}
