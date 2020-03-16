package Solution

import "fmt"

//
func longestPalindrome1(s string) string {
	if len(s) == 1 {
		return s
	}
	ans := ""
	ansLength := 0

	for i := 0; i < len(s); i ++ {
		for j:= len(s); j > i + ansLength; j-- {
			substr := s[i:j]
			if isPalindrome(substr) && len(substr) > ansLength {
				ans = substr
				ansLength = len(substr)
			}
		}
	}
	return ans
}

func isPalindrome(s string) bool {
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

func longestPalindrome2(s string) string {
	if len(s) == 0 {
		return ""
	}
	start, end := 0,0
	for i :=0; i < len(s); i ++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		len := Max(len1, len2)
		if len > end-start {
			start = i - (len-1) / 2
			end = i + len/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	L, R := left, right
	for L >= 0 && R < len(s) && s[L]==s[R] {
		L--
		R++
	}
	return R - L - 1
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func longestPalindrome3( s string) string {
	return longestPalindromeLinear(s)
}

func initManacherStr(s string) string {
	ans : = make([]rune, 0)
	ans = append(ans, '$')
	for i := 0; i < len(s); i ++ {
		ans = append(ans, rune(s[i]), '$')
	}
	return string(ans)
}

func longestPalindromeLinear( in string) string {
	//
	s := initManacherStr(in)
	c, max := 0, 0

	//
	P := make([]int, len(s))
	//
	for i := 1; i < len(s) - 1; i ++ {
		i_mirror := 2*c - i
		if max > i {
			P[i] = Min(P[i_mirror], max-i)

		} else {
			P[i] = 0
		}
		for (i+P[i]+1) < len(s) && (i-P[i]-1) >= 0 && 
			s[i+P[i]+1] == s[i-P[i]-1] {
				P[i]++
		}
		if i+P[i] > max {
			c = i
			max = i + P[i]
		}
	}
	return extractLongest(in, P)
}

func extractLongest(s string, P []int) string {
	longestCenter, longestLength := 0, 0
	for i, v := range P {
		if v > longestLength {
			longestLength = v
			longestCenter = i
		}
	}
	offset := (longestCenter - longestLength) / 2
	return s[offset: offset+longestLength]
}

