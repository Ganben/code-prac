package solution

func canJump(nums []int) bool {
	maxreach := 0
	for i, e := range nums {
		if maxreach < i {
			break
		}
		if maxreach < i+e {
			maxreach = i + e
		}
	}
	if maxreach < len(nums)-1 {
		return false
	}
	return true
}
