package Solution

import (
	"reflect"
	"testing"
)

var cases = []struct {
	name string
	inputs int
	expect string
}{
	{"1 test 1", 3, "III"},
	{"2 test 2", 4, "IV"},
	{"3 test 3", 19, "XIX"},
	{"4 test 4", 58, "LVIII"},
	{"5 test 5", 1994, "MCMXCIV"},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := intToRoman(c.inputs)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v got %v", c.expect, ret)
			}
		})
	}
}