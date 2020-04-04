package Solution

import (
	"reflect"
	"testing"
)

var cases = []struct {
	name string
	inputs [][]int
	expect [][]int
}{
	{
		"case 1",
		[][]int{{10, 1, 2, 7, 6, 1, 5}, {8}},
		[][]int{{1, 7}, {1, 2, 5}, {2, 6}, {1, 1, 6}},
	},
	{
		" case 2",
		[][]int{{2, 5, 2, 1, 2}, {5}},
		[][]int{{1, 2, 2}, {5}},
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := combinationSum2(c.inputs[0], c.inputs[1][0])
			if !isEqual(got, c.expect) {
				t.Fatalf("expect %v got %v", c.expect, got)
			}
		})
	}
}

func isEqual(x, y [][]int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := 0; i < len(y); i ++ {
		flag := false
		for j := 0; j < len(y); j ++ {
			if reflect.DeepEqual(y[i], x[j]) {
				flag = true
			}
		}
		if flag == false {
			return false
		}
	}
	return true
}