package solution

import (
	"reflect"
	"testing"
)

var cases = []struct {
	name   string
	input1 int
	input2 int
	expect string
}{
	{
		"test 1",
		3,
		3,
		"213",
	},
	{
		"test 2",
		9,
		278621,
		"792861534",
	},
	{
		"test 3",
		2,
		2,
		"21",
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		got := getPermutation(c.input1, c.input2)
		if !reflect.DeepEqual(got, c.expect) {
			t.Fatalf("case %s incorrect by %v", c.name, got)
		}
	}
}
