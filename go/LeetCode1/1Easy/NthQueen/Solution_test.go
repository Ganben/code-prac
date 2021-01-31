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
	{
		"case 5 queen",
		5,
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
		t.Run(c.name, func(t *testing.T) {
			got := solveNQueens(c.input)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("case %s failed with %v", c.name, got)
			}
		})
	}
}

func TestTotaoNQueen(t *testing.T) {
	t.Run("4", func(t *testing.T) {
		got := totalNQueens(4)
		if got != 2 {
			t.Fatalf("incorrect")
		}
	})
}
