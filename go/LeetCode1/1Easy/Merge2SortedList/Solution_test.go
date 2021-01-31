package Solution

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		data1 := []int{1, 2, 4}
		data2 := []int{1, 3, 5}

		got := mergeTwoLists(MakeListNotes(data1), MakeListNotes(data2))
		want := MakeListNotes([]int{1, 1, 2, 3, 4, 4})

		if !Equal(got, want) {
			t.Fatalf("expect not")
		}
	})

}

func Equal(x, y *ListNode) bool {
	for x == nil || y == nil {
		if x == nil && y != nil {
			fmt.Println(x, y)
			return false
		}
		if x != nil && y == nil {
			fmt.Println(x, y)
			return false
		}
		if x.Val != y.Val {
			fmt.Println(x, y)
			return false
		}
		x = x.Next
		y = y.Next
	}
	return true
}