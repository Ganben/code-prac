package diagonaltrav

import "testing"

func TestSolve(t *testing.T) {
	mat := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	res := findDiagonalOrder(mat)
	if len(res) != 9 || res[8] != 9 {
		t.Log(res)
		t.FailNow()
	}
	t.Log(res)
}
