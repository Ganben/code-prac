package main

import "fmt"

func longestPalindrome(s string) int {
	hMap := map[byte]int{}
	for i := 0; i < len(s); i++ {
		hMap[byte(s[i])]++
	}
	fmt.Printf("%v\n", hMap)
	ans := 0
	for _, v := range hMap {
		ans += v / 2 * 2
		if v%2 == 1 && ans%2 == 0 {
			ans++
		}
	}
	return ans

}

func main() {
	fmt.Printf("ans:%v\n", longestPalindrome("abccccdd"))
}
