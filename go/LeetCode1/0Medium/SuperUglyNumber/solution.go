package SuperUglyNumber

import (
	"fmt"
	"sort"
)

func nthSuperUglyNumber(n int, primes []int) int {
	multiplier := make([]int, len(primes))
	for i, _ := range multiplier {
		multiplier[i] = 1
	}
	if n == 1 {
		return 1
	}
	ans := 1
	for f := 2; f <= n; f++ {
		multiplier, ans = testSmallest(primes, multiplier, ans)
	}
	return ans
}

func testSmallest(prs, mst []int, cur int) ([]int, int) {
	results := make([]int, len(prs))
	ret := mst
	for i, _ := range results {
		results[i] = prs[i] * mst[i]
	}
	b2 := 0
	minCur := cur
	toSortArray := results
	toSortArray = sort.IntSlice(toSortArray)
	// find min next
	if toSortArray[0] > cur {
		minCur = toSortArray[0]
	} else {
		for _, v := range toSortArray {
			if v > minCur {
				minCur = v
				break
			}
		}
	}

	for i, v := range results {
		if v == minCur {
			b2 = v
			fmt.Printf("%d\n", i)
			if ret[i] == 1 {
				ret[i] = prs[0]
			} else {
				for i, v := range prs {
					if v == ret[i] {
						if i < len(prs)-1 {
							ret[i] = prs[i+1]
						} else {
							break
						}
					}
				}
			}

			break
		}
	}

	return ret, b2
}
