package solution

import (
	"testing"
	"reflect"
)

var cases = []struct {
	name string
	inputs []string
	expect []string
}{
	{
		"test 1",
		[]string{
			"This",
			"is",
			"an",
			"example",
			"of",
			"text",
			"justification.",
		},
		[]string{
			"This    is    an",
			"example  of text",
			"justification.  ",
		},
	},
	{
		"test 2",
		[]string{
			"What",
			"must",
			"be",
			"acknowledgment",
			"shall",
			"be",
		},
		[]string{
			"What   must   be",
			"acknowledgment  ",
			"shall be        ",
		},
	},
}

func TestSolution(t *testing.T) {
for _, c := range cases {
	got := fullJustify(c.inputs, 16)
	if !reflect.DeepEqual(c.expect, got) {
		t.Fatalf("%s case failed by %v", c.name, got)
	}
}
}