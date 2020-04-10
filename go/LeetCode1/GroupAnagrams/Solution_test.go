package Solution

import (
	"reflect"
	"testing"
)

var cases = []struct {
	name   string
	inputs []string
	expect [][]string
}{
	{
		"test 1",
		[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
		[][]string{
			{"ate", "eat", "tea"},
			{"nat", "tan"},
			{"bat"},
		},
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		got := groupAnagrams(c.inputs)
		if !reflect.DeepEqual(got, c.expect) {
			t.Fatalf("%s incorrect with %v", c.name, got)
		}
	}
}
