package Solution

import (
	"fmt"
	"testing"
)

var cases = []struct {
	name string
	input1 []int
	input2 int
	expect [][]int
}{
	{
		"1 test 1",
		[]int{1, 0, -1, 0, -2, 2},
		0,
		[][]int{
			{-1, 0, 0, 1},
			{-2, -1, 1, 2},
			{-2, 0, 0, 2},
		},
	},
}

func TestSolution(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := fourSum(c.input1, c.input2)
			fmt.Println(got)
		})
	}
}

func TestSolution1(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := fourSum1(c.input1, c.input2)
			fmt.Println(got)
		})
	}
}