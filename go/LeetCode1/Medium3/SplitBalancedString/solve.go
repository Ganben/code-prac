package splitbalancedstring

func balancedStringSplit(s string) int {
	cnt := 0
	flag := 0
	for _, v := range s {
		if v == 'R' {
			flag += 1
		} else {
			flag += -1
		}
		if flag == 0 {
			cnt += 1
		}

	}
	return cnt
}
