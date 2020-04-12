package Solution

import (
	"reflect"
	"testing"
)

var cases = []struct {
	name   string
	input  int
	expect [][]string
}{
	{
		"case 4 queen",
		4,
		[][]string{
			{
				".Q..",
				"...Q",
				"Q...",
				"..Q.",
			},
			{
				"..Q.",
				"Q...",
				"...Q",
				".Q..",
			},
		},
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		got := solveNQueens(c.input)
		if !reflect.DeepEqual(got, c.expect) {
			t.Fatalf("case %s failed with %v", c.name, c.expect)
		}
	}
}
