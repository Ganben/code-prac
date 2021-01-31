package Solution

import (
	"reflect"
	"testing"
)

var cases = []struct{
	name string
	inputs [][]int
	expect [][]int
}{
	{
		"test case 1",
		[][]int{{2, 3, 6, 7}, {7}},
		[][]int{{7}, {2, 2, 3}},
	},
	{
		"test case 2",
		[][]int{{2, 3, 5}, {8}},
		[][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
	},
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

func TestSolution(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := combinationSum(c.inputs[0], c.inputs[1][0])
			if !isEqual(got, c.expect) {
				t.Fatalf("incorrect result compare")
			}
		})
	}
}