package ExcelColTitle

func convertToTitle(n int) string {
	ans := ""
	s := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	t, r := 0, 0
	for i := 0; i < n; i++ {

		t = i % 26
		r = i / 26
		if r == 0 {
			ans += string(s[t])
		} else {
			ans += string(s[r-1])
			ans += string(s[t-1])
		}

	}
	return ans
}
