package solution

import (
	"reflect"
	"testing"
)

var cases = []struct {
	name   string
	inputs []string
	mwidth int
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
		16,
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
		16,
		[]string{
			"What   must   be",
			"acknowledgment  ",
			"shall be        ",
		},
	},
	{
		"test 3",
		[]string{
			"Listen",
			"to",
			"many,",
			"speak",
			"to",
			"a",
			"few.",
		},
		6,
		[]string{
			"Listen",
			"to    ",
			"many, ",
			"speak ",
			"to   a",
			"few.  ",
		},
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		got := fullJustify(c.inputs, c.mwidth)
		if !reflect.DeepEqual(c.expect, got) {
			t.Fatalf("%s case failed by %v", c.name, got)
		}
	}
}
