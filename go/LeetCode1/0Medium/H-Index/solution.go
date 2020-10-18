package Index

import "sort"

func hIndex(citations []int) int {
	if len(citations) == 0 {
		return 0
	}
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))
	dp := make([]int, citations[0]+1)
	i := 0
	for i < len(citations) {
		num := citations[i]
		if num != citations[0] {
			dp[num] = dp[num+1]
		} else {
			dp[num] = 0
		}
		if dp[num] == num {
			return num
		}
		for i < len(citations) && citations[i] == num {
			dp[num]++
			i++
			if dp[num] == num {
				return num
			}
		}
		down := 0
		if i >= len(citations) {
			down = 0
		} else {
			down = citations[i] + 1
		}
		for j := down; j < num; j++ {
			dp[j] = dp[num]
			if dp[j] == j {
				return j
			}
		}

	}
	return 0
}
