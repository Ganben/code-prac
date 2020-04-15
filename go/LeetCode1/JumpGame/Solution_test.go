package solution

import "testing"

var cases = []struct {
	name   string
	inputs []int
	expect bool
}{
	{
		"test 1",
		[]int{2, 3, 1, 1, 4},
		true,
	},
}

func TestJumpGame(t *testing.T) {
	for _, c := range cases {
		got := canJump(c.inputs)
		if got != c.expect {
			t.Fatalf("%s case incorrect by %v", c.name, got)
		}
	}
}
