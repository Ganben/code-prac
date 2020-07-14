package ReverseWords

import "strings"

func reverseWords(s string) string {
	sArr := strings.Split(s, " ")
	for i, v := range sArr {
		sArr[i] = reverseW(v)
	}
	ans := ""
	for i, v := range sArr {
		ans += v
		if i < len(sArr)-1 {
			ans += " "
		}
	}
	return ans
}

func reverseW(s string) string {
	ans := ""
	for i := 0; i < len(s); i++ {
		ans += string(s[len(s)-1-i])
	}
	return ans
}
