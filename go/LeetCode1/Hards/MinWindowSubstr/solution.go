package MinWindowSubstr

func minWindow(s string, t string) string {
	left, right := 0, 0
	flag := 0
	shortest := len(s)
	ans := ""
	for left < len(s) {
		// slide right
		if flag == 0 {
			right++
		} else {
			left++
		}
		if left == len(s) || right == len(s) {
			break
		}
		// test substr
		r := testSubstr(s[left:right+1], t)
		if r {
			if right-left < shortest {
				ans = s[left : right+1]
				shortest = len(ans)
			}
			flag = 1
		} else {
			flag = 0
		}

	}
	return ans
}

func testSubstr(s, t string) bool {
	count := make([]int, len(t))
	for _, c := range s {
		for i, k := range t {
			if c == k {
				count[i] = 1
			}
		}
	}
	for _, v := range count {
		if v == 0 {
			return false
		}
	}
	return true
}
