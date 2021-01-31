package Solution

import ( 
	"reflect"
	"testing"
)

var cases = []struct {
	name string
	input1 string
	input2 int
	expect string
}{
	{"1 test 1", "PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
	{"2 test 2", "PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
}


func TestSolution(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := convert(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected %v got %v", c.expect, ret)
			}
		})
	}
}

func TestSolution2(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := convert2(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected %v got %v", c.expect, ret)
			}
		})
	}
}