package main

import "fmt"

func findTheDifference(s string, t string) byte {
	hmap := map[byte]int{}
	for i := 0; i < len(t); i++ {
		_, has := hmap[t[i]]
		if !has {
			hmap[t[i]] = 1
		} else {
			hmap[t[i]] += 1
		}
	}
	for i := 0; i < len(s); i++ {
		hmap[s[i]] -= 1
	}
	for k, v := range hmap {
		if v == 1 {
			return k
		}
	}
	return t[0]
}

func main() {
	s1 := "abcd"
	s2 := "abcde"
	fmt.Printf("%v\n", findTheDifference(s1, s2))
}
