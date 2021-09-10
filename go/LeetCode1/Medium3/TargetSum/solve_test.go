package targetsum

import "testing"

func TestCase(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1}
	target := 3
	res := findTargetSumWays(nums, target)
	if res != 5 {
		t.Log(res)
		t.FailNow()
	}
}
