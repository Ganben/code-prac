package Solution

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	if candidates == nil || len(candidates) == 0 {
		return result
	}
	sort.Ints(candidates)
	aset := []int{}
	explore(candidates, target, 0, &result, aset)
	return result
}

func explore(candidates []int, target int, pos int, result *[][]int, aset []int) {
	if target < 0 {
		return
	}
	if target == 0 {
		c := make([]int, len(aset))
		copy(c, aset)
		*result = append(*result, c)
		return
	}

	for i := pos; i < len(candidates); i ++ {
		aset = append(aset, candidates[i])
		explore(candidates, target - candidates[i], i, result, aset)
		aset = aset[:len(aset) - 1]
	}
}