package solution

func grayCode(n int) []int {
	ans, head := []int{0}, 1
	for i := 0; i < n; i++ {
		for j := len(ans) - 1; j > -1; j-- {
			ans = append(ans, head+ans[j])

		}
		head <<= 1
	}
	return ans
}
