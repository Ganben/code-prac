package Solution

import (
	"reflect"
	"testing"
)

var cases = []struct{
	name string
	inputs string
	expect int
}{
	{"1 test 1", "abcabcbb!", 3},
	{"2 test 2", "bbbbb", 1},
	{"3 test 3", "pwwkew", 3},
}

func TestSolution(t *testing.T) {
	//
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T){
			ret := lengthOfLongestSubstring(c.inputs)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v got %v", c.expect, ret)
			}
		})
	}

}