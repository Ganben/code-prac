package Solution

import (
	"reflect"
	"testing"
)

var cases = []struct {
	name string
	inputs []string
	expect bool
}{
	{"test 1", []string{"aa", "a"}, false},
	{"test 2", []string{"aa", "a*"}, true},
	{"test 3", []string{"ab", ".*"}, true},
	{"test 4", []string{"aab", "c*a*b"}, true},
	{"test 5", []string{"mississippi", "mis*is*p*."}, false},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := isMatch(c.inputs[0], c.inputs[1])
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected %v got %v", c.expect, ret)
			}
		})
	}
}