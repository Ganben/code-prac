package main

func countSegments(s string) int {
	flag := 0
	cnt := 0
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 && rune(s[i]) != rune(' ') {
			if flag == 1 {
				cnt++
				break
			}
			//any? direct
			cnt++
			break
		}
		if rune(s[i]) != rune(' ') {
			flag = 1
		}
		if rune(s[i]) == rune(' ') && flag == 1 {
			flag = 0
			cnt++
		}

	}
	return cnt
}
