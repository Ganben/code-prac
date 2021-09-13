package nextgreaterelement

import "testing"

func TestSolve(t *testing.T) {
	nums1 := []int{2, 4}
	nums2 := []int{1, 2, 3, 4}
	res := nextGreaterElement(nums1, nums2)
	if res[0] != 3 && res[1] != -1 {
		t.Log(res)
		t.FailNow()
	}
}
