package solution

import "testing"

var cases = []struct {
	name   string
	inputs [][]byte
	expect int
}{
	{
		"test 1",
		[][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'},
		},
		1,
	},
	{
		"test 2",
		[][]byte{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'},
		},
		3,
	},
	{
		"test 3",
		[][]byte{
			{'1', '1', '1'},
			{'0', '1', '0'},
			{'1', '1', '1'},
		},
		1,
	},
	{
		"test 4",
		[][]byte{
			{'1', '0', '1', '1', '1'},
			{'1', '0', '1', '0', '1'},
			{'1', '1', '1', '0', '1'},
		},
		1,
	},
}

func TestNumIsland(t *testing.T) {
	for _, c := range cases {
		got := numIslands(c.inputs)
		if got != c.expect {
			t.Fatalf("%s incorrect by %v", c.name, got)
		}
	}
}
