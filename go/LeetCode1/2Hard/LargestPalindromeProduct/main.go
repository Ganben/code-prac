package main

import (
	"math"
	"strconv"
)

func largestPalindrome(n int) int {
	if n == 1 {
		return 9
	}
	//
	max := math.Pow10(n) - 1
	// find max palindrom
	for i := max - 1; i > max/10; i-- {
		s1 := strconv.Itoa(int(i))
		s2 := reverse([]rune(s1))
		rev, _ := strconv.Atoi(s1 + s2)
		//
		for x := int(max); x*x >= rev; x-- {
			if rev%x == 0 {
				return rev % 1337
			}
		}
	}
	return -1
}

func reverse(s []rune) string {
	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	return string(s)
}
