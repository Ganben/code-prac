package Solution

import (
	"reflect"
	"testing"
)

var cases = []struct {
	name string
	inputs string
	expect int
}{
	{
		"test 1",
		"(()",
		2,
	},
	{
		"test 2",
		")()())",
		4,
	},
	{
		"test 3",
		"()(())",
		6,
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := longestValidParentheses(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expect %v got %v", c.expect, got)
			}
		})
	}
}

func TestSolution2(t *testing.T) {
	for _, c := range cases {
		got := longestValidParentheses2(c.inputs)
		if !reflect.DeepEqual(got, c.expect) {
			t.Fatalf("expect %v got %v", c.expect, got)
		}
	}
}