package Solution

import (
	"reflect"
	// "strconv"
	"testing"
)

var cases = []struct {
	name string
	inputs int
	expect string
}{
	{
		"test case 1",
		1,
		"1",
	},
	{
		"test case 2",
		4,
		"1211",
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := countAndSay(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expect %v got %v", c.expect, got)
			}
		})
	}
}