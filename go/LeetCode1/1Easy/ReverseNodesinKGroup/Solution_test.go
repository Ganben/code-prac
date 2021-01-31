package Solution

import "testing"

var cases = []struct {
	name string
	input1 *ListNode
	input2 int
	expect *ListNode
}{
	{
		"testcase 1",
		UnmarshalListBySlice([]int{1,2,3,4,5}),
		2,
		UnmarshalListBySlice([]int{2,1,4,3,5}),
	},
	{
		"test case 2",
		UnmarshalListBySlice([]int{1,2,3,4,5}),
		3,
		UnmarshalListBySlice([]int{3,2,1,4,5}),
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		got := reverseKGroup(c.input1, c.input2)
		if !isEqual(got, c.expect) {
			t.Fatalf("case fail %s", c.name)
		}
	}
}