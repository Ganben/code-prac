package solution

import "testing"

var cases = []struct {
	name   string
	input1 string
	input2 int
	input3 string
	input4 int
	expect int
}{
	{
		"test 1",
		"acb",
		4,
		"ab",
		2,
		2,
	},
	{
		"test 2",
		"acb",
		1,
		"acb",
		1,
		1,
	},
	{
		" test 3",
		"aaa",
		3,
		"aa",
		1,
		4,
	},
	{
		"test 4",
		"adasfdgftadwfwfereredsfssas",
		500,
		"fw",
		555,
		1,
	},
}

func TestMaxRepe(t *testing.T) {
	for _, c := range cases {
		got := getMaxRepetitions(c.input1, c.input2, c.input3, c.input4)
		if got != c.expect {
			t.Fatalf("%s incorrect by %v", c.name, got)
		}
	}
}
