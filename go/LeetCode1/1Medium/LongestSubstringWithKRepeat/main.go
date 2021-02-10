package main

import (
	"fmt"
	"math"
	"strings"
)

func longestSubstring(s string, k int) int {
	if len(s) < k {
		return 0
	}
	hmap := map[rune]int{}
	mintimes := math.MaxInt32
	for _, ch := range s {
		_, exist := hmap[ch]
		if exist {
			hmap[ch]++
		} else {
			hmap[ch] = 1
		}
	}
	minc := ""
	for key, val := range hmap {
		if val < mintimes {
			mintimes = val
			minc = string(key)
		}
	}
	if mintimes >= k {
		return len(s)
	}
	sarr := []string{}
	sarr = strings.Split(s, minc)
	maxtimes := 1 - math.MaxInt32
	for _, sarrs := range sarr {
		t := longestSubstring(sarrs, k)
		if t > maxtimes {
			maxtimes = t
		}
	}
	return maxtimes
}

func main() {
	s := "ababbc"
	fmt.Printf("%v\n", longestSubstring(s, 2))
}
