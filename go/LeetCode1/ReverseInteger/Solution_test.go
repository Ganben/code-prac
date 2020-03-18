package Solution

import (
	"reflect"
	"testing"
)

var cases = []struct {
	name string
	inputs int
	expect int
}{
	{"test case 1 ", 123, 321},
	{"test case 2", -123, -321},
	{"test case 3", 210, 12},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := reverse(c.inputs)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected %v got %v", c.expect, ret)
			}
		})
	}
}