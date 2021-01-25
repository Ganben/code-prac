package main

import "fmt"

func numEquivDominoPairs(dominoes [][]int) int {
	ans := 0
	if len(dominoes) < 2 {
		return 0
	}
	for i := 0; i < len(dominoes); i++ {
		for j := i + 1; j < len(dominoes); j++ {
			if arrEqual(dominoes[i], dominoes[j]) {
				ans++
			}
		}
	}
	return ans
}

func numEquivDominoPairs2(dominoes [][]int) int {
	ans := 0
	cnt := [100]int{}
	for _, d := range dominoes {
		if d[0] > d[1] {
			d[0], d[1] = d[1], d[0]
		}
		v := d[0]*10 + d[1]
		ans += cnt[v]
		cnt[v]++
	}
	return ans

}

func arrEqual(ar1, ar2 []int) bool {
	if ar1[0] == ar2[0] && ar1[1] == ar2[1] {
		return true
	}
	if ar1[0] == ar2[1] && ar1[1] == ar2[0] {
		return true
	}
	return false
}

func main() {
	dominoes := [][]int{
		{1, 2},
		{2, 1},
		{3, 4},
		{5, 6},
	}
	fmt.Printf("%v", numEquivDominoPairs(dominoes))
}
