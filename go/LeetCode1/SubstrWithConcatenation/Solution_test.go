package Solution

import (
	"reflect"
	"testing"
)

var cases = []struct {
	name string
	input1 string
	input2 []string
	expect []int
}{
	{
		"testcase 1",
		"barfoothefoobarman",
		[]string{"foo", "bar"},
		[]int{0, 9},
	},
	{
		"testcase 2",
		"wordgoodstudentgoodword",
		[]string{"word", "student"},
		[]int{},
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := findSubstring(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect %v got %v", c.expect, ret)
			}
		})
	}
}