package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	if len(magazine) < len(ransomNote) {
		return false
	}

	//
	p2 := 0
MainLoop:
	for i := 0; i < len(ransomNote); i++ {
		for p2 < len(magazine) {
			if ransomNote[i] == magazine[p2] {
				p2++
				continue MainLoop
			}
			return false
		}

	}

	return true
}

func main() {
	r1, r2 := "a", "b"
	fmt.Printf("%v", canConstruct(r1, r2))
}
