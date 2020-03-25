package Solution

import (
	"testing"
	"reflect"
)

var cases = []struct {
	name string
	inputs int
	expect []string
}{
	{"1 test 1", 3, []string{
		"((()))",
		"(()())",
		"(())()",
		"()(())",
		"()()()",
	}},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := generateParenthesis(c.inputs)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect %v got %v", c.expect, ret)
			}
		})
	}
}