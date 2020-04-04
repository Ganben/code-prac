package Solution

func firstMissingPositive(nums []int) int {
	i := 1
	for _, c := range nums {
		if c > 0 && c == i {
			i += 1
		}
	}
	return i
}