package main

import (
	"fmt"
	"math"
)

type Index struct {
	index int
	count int
}

func firstUniqChar(s string) int {

	hashMap := map[rune]*Index{}
	for i, v := range s {
		_, exist := hashMap[v]
		if exist {
			hashMap[v].count++
		} else {
			hashMap[v] = &Index{i, 1}
		}
	}
	var p int
	p = math.MaxInt16
	for _, value := range hashMap {

		if value.count == 1 {
			if p > value.index {
				p = value.index
			}
		}

	}
	if p < math.MaxInt16 {
		return p
	}
	return -1

}

func main() {
	s := "leetcode"
	fmt.Printf("%v", firstUniqChar((s)))
	s2 := "loveleetcode"
	fmt.Printf("%v", firstUniqChar((s2)))
}
