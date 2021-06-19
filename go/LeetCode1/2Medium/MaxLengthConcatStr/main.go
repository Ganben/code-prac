package main

import "math/bits"

func maxLength(arr []string) int {
	var ans int
	masks := []int{}
outer:
	for _, s := range arr {
		mask := 0
		for _, ch := range s {
			ch -= 'a'
			// test for duplicated char
			if mask>>ch&1 > 0 {
				continue outer
			}
			mask |= 1 << ch
		}
		masks = append(masks, mask)
	}
	var backtrack func(int, int)
	backtrack = func(pos, mask int) {
		if pos == len(masks) {
			ans = max(ans, bits.OnesCount(uint(mask)))
			return
		}
		if mask&masks[pos] == 0 {
			backtrack(pos+1, mask|masks[pos])
		}
		backtrack(pos+1, mask)
	}
	backtrack(0, 0)
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
