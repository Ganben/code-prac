package Solution

import (
	"reflect"
	"testing"
)

var cases = []struct {
	name string
	inputs string
	expect string
}{
	{"1 test 1", "badab", "bad"},
	{"2 test 2", "dbbd", "bb"},
	{"3 test 3", "a", "a"},
	{"4 test 4", "bb", "bb"},
}

func TestSolution1(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := longestPalindrome1(c.inputs)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected %v, got %v", c.expect, ret)
			}
		})
	}	
}

func TestSolution3(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := longestPalindrome3(c.inputs)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect %v got %v", c.expect, ret)
			}
		})
	}
}
