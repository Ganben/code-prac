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
	{"1 test 1", "42", 42},
	{"2 test 2", "-42", -42},
	{"3 test 3", "4193 with", 4193},
}

func TestSolution( t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := myAtoi(c.inputs)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect %v got %v",
			c.expect, ret)
			}
		})
	}
}