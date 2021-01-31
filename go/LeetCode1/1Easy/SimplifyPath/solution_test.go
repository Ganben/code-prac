package SimplifyPath

import (
	"reflect"
	"testing"
)

var cases = []struct {
	name   string
	inputs string
	expect string
}{
	{
		"test 1",
		"/home/",
		"/home",
	},
	{
		"test 2",
		"/../",
		"/",
	},
	{
		"test 3",
		"/home//foo/",
		"/home/foo",
	},
	{
		"test 4",
		"/a/./b/../../c/",
		"/c",
	},
	{
		"test 5",
		"/a/../../b/../c//.//",
		"/c",
	},
	{
		"test 6",
		"/a//b////c/d//././/..",
		"/a/b/c",
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		got := simplifyPath(c.inputs)
		if !reflect.DeepEqual(got, c.expect) {
			t.Fatalf("%s incorrect by %s", c.name, got)
		}
	}
}
