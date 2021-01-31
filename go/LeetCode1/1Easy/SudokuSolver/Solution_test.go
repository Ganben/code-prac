package Solution

import (
	"reflect"
	"testing"
)

var cases = []struct {
	name string
	inputs [][]byte
	expect [][]byte
}{
	{
		"test case 1",
		[][]byte{
			[]byte("..9748..."),
			[]byte("7........"),
			[]byte(".2.1.9..."),
			[]byte("..7...24."),
			[]byte(".64.1.59."),
			[]byte(".98...3.."),
			[]byte("...8.3.2."),
			[]byte("........6"),
			[]byte("...2759.."),
		},
		[][]byte{
			[]byte("519748632"),
			[]byte("783652419"),
			[]byte("426139875"),
			[]byte("357986241"),
			[]byte("264317598"),
			[]byte("198524367"),
			[]byte("975863124"),
			[]byte("832491756"),
			[]byte("641275983"),
		},
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			solveSudoku(c.inputs)
			if !reflect.DeepEqual(c.inputs, c.expect) {
				t.Fatalf("incorrect")
			}
		})
	}
}