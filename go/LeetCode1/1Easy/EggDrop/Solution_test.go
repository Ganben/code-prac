package Solution

import "testing"

var cases = []struct {
	name   string
	inputK int
	inputN int
	expect int
}{
	{
		"case 1",
		1,
		2,
		2,
	},
	{
		"case 2",
		2,
		6,
		3,
	},
	{
		"case 3",
		3,
		14,
		4,
	},
	{
		"case 4",
		2,
		2,
		2,
	},
	{
		"case 5",
		2,
		1,
		1,
	},
	{
		"case 6",
		2,
		4,
		3,
	},
	{
		"case 6.4",
		1,
		4,
		4,
	},
	{
		"case 7",
		2,
		9,
		4,
	},
	{
		"case 8",
		2,
		14,
		5,
	},
	{
		"case 9",
		2,
		1000,
		45,
	},
	{
		"case 10",
		4,
		8000,
		22,
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := superEggDrop(c.inputK, c.inputN)
			if got != c.expect {
				t.Fatalf("%s failed with %v", c.name, got)
			}
		})

	}
}
