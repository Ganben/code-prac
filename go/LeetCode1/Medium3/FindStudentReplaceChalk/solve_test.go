package findstudentreplacechalk

import "testing"

func TestSolve(t *testing.T) {
	chalk := []int{5, 1, 5}
	k := 22
	res := chalkReplacer(chalk, k)
	if res != 0 {
		t.Log(res)
		t.FailNow()
	}
}
