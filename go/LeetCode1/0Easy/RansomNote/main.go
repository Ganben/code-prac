package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	if len(magazine) < len(ransomNote) {
		return false
	}

	//
	p2 := 0
	acc := 0
	vis := make([]bool, len(magazine))
	for i := 0; i < len(vis); i++ {
		vis[i] = false
	}
MainLoop:
	for i := 0; i < len(ransomNote); i++ {
		for acc < len(ransomNote)*len(magazine) {
			if p2 >= len(magazine) {
				p2 -= len(magazine)
			}
			fmt.Printf("%v,%v,%v\n", i, p2, vis)
			if ransomNote[i] == magazine[p2] && !vis[p2] {
				vis[p2] = true
				p2++
				continue MainLoop
			}
			p2++
			acc++

		}
		return false
	}

	return true
}

func main() {
	r1, r2 := "bcjefgecda", "hfebdiicigfjahdddiahdajhaidbdgjihdbhgfbbccfdfggdcacccaebh"
	fmt.Printf("%v", canConstruct(r1, r2))
}
