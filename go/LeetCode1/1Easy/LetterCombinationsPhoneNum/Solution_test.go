package Solution

import (
	"reflect"
	"testing"
)

var cases = []struct {
	name string
	inputs string
	expect []string
}{
	{"1 test 1", "23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := letterCombinations(c.inputs)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected %v got %v", c.expect, ret)
			}
		})
	}

}

func TestSolution2(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := depthSolution(c.inputs)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected %v got %v", c.expect, ret)
			}
		})
	}
}

func TestSolution3(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := breadthSolution(c.inputs)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected %v got %v", c.expect, ret)
			}
		})
	}
}