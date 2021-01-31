package EditDistance

import "testing"

var cases = []struct {
	name   string
	input1 string
	input2 string
	expect int
}{
	{
		"test 1",
		"horse",
		"ros",
		3,
	},
	{
		"test 2",
		"intention",
		"execution",
		5,
	},
}

func TestMindistance(t *testing.T) {
	for _, c := range cases {
		got := minDistance(c.input1, c.input2)
		if got != c.expect {
			t.Fatalf("%s incorrect by %v", c.name, got)
		}
	}
}
