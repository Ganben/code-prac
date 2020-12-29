package main

import "fmt"

func reverseVowels(s string) string {
	indexArr := []int{}
	ss := []rune{}
	for i, v := range s {
		if isVowel(v) {
			indexArr = append(indexArr, i)
		}
		ss = append(ss, v)
	}
	for i := 0; i < (len(indexArr) / 2); i++ {
		ss[indexArr[i]], ss[indexArr[len(indexArr)-1-i]] = ss[indexArr[len(indexArr)-1-i]], ss[indexArr[i]]
	}
	ans := ""
	for _, v := range ss {
		ans += string(v)
	}
	return ans

}

func isVowel(c rune) bool {
	if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
		return true
	}
	if c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' {
		return true
	}
	return false
}

func main() {
	strInput := "hello"
	fmt.Printf("%v", reverseVowels((strInput)))
}
