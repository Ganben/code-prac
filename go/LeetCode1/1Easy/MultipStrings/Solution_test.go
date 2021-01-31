package Solution

import (
	"testing"
	"reflect"
)

var cases = []struct {
	name string
	inputs []string
	expect string
}{
	{
		"test 1",
		[]string{"2", "3"},
		"6",
	},
	{
		"test 2",
		[]string{"123", "456"},
		"56088",
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		got := multiply(c.inputs[0], c.inputs[1])
		if !reflect.DeepEqual(got, c.expect) {
			t.Fatalf("expect %v got %v", c.expect, got)
		}
	}
}