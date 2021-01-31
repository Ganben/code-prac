package solution

import "testing"

var cases = []struct {
	name   string
	input1 int
	input2 int
	expect int
}{
	{
		"test 1",
		3,
		2,
		3,
	},
	{
		"test 2",
		7,
		3,
		28,
	},
	{
		"test 3",
		23,
		12,
		193536720,
	},
	{
		"test 4",
		13,
		23,
		548354040,
	},
	{
		"test 5",
		11,
		18,
		8436285,
	},
	{
		"test 6",
		2,
		100,
		100,
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		got := uniquePathsB(c.input1, c.input2)
		if got != c.expect {
			t.Fatalf("%s incorrect by %v", c.name, got)
		}
	}
}

func TestSolution2(t *testing.T) {
	for _, c := range cases {
		got := uniquePaths(c.input1, c.input2)
		if got != c.expect {
			t.Fatalf("%s incorrect by %v", c.name, got)
		}
	}
}
