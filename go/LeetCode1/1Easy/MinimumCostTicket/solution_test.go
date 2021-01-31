package MinimumCostTicket

import "testing"

var cases = []struct {
	name   string
	input1 []int
	input2 []int
	expect int
}{
	{
		"test 1",
		[]int{1, 4, 6, 7, 8, 20},
		[]int{2, 7, 15},
		11,
	},
	{
		"test 2",
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31},
		[]int{2, 7, 15},
		17,
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		got := mincostTickets(c.input1, c.input2)
		if got != c.expect {
			t.Fatalf("%s incorrect by %v", c.name, got)
		}
	}
}
