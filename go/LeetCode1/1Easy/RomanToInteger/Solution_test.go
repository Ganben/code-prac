package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

var cases = []struct {
	name string
	inputs string
	expect int
}{
	{"test case 1", "III", 3},
	{"test case 2", "IV", 4},
	{"test case 3", "XI", 11},
	{"test case 4", "LIV", 54},
	{"test case 5", "MCMXCIV", 1994},
}

func TestSolution(t *testing.T) {
	//
	for i, c := range cases {
		t.Run(c.name + " " + strconv.Itoa(i), func(t *testing.T) {
			got := romanToInt(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected %v got %v", c.expect, got)
			}
		})
	}

}