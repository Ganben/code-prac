package main

import (
	"fmt"
	"math/bits"
)

func findNumOfValidWords(words []string, puzzles []string) []int {
	const puzzleLength = 7
	cnt := map[int]int{}
	for _, s := range words {
		mask := 0
		for _, ch := range s {
			mask |= 1 << (ch - 'a')
		}
		if bits.OnesCount(uint(mask)) <= puzzleLength {
			cnt[mask]++
		}
	}

	ans := make([]int, len(puzzles))
	for i, s := range puzzles {
		first := 1 << (s[0] - 'a')
		// enum
		mask := 0
		for _, ch := range s[1:] {
			mask |= 1 << (ch - 'a')

		}
		subset := mask
		for {
			ans[i] += cnt[subset|first]
			subset = (subset - 1) & mask
			if subset == mask {
				break
			}
		}

	}
	return ans
}

func main() {
	words := []string{
		"aaaa",
		"asas",
		"able",
		"ability",
		"actor",
		"access",
	}
	puzzles := []string{
		"aboveyz",
		"abrodyz",
		"abslute",
		"absoryz",
		"actresz",
		"gaswxyz",
	}
	fmt.Printf("%v\n", findNumOfValidWords(words, puzzles))
}
