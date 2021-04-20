package main

import "fmt"

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		flag := true
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				flag = false
				break
			}

		}
		if flag {
			return i
		}
	}
	//
	return -1
}

func main() {
	haystack := "hello"
	needle := "ll"
	fmt.Printf("%v\n", strStr(haystack, needle))
}
