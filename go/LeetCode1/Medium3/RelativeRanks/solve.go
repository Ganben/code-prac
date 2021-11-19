package relativeranks

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	res := []string{}
	length := len(score)
	for i := 0; i < length; i++ {
		score[i] = score[i]*100000 + i
	}
	sort.Ints(score)
	res = make([]string, length)
	for i := 0; i < length; i++ {
		// 分数大的名次高
		switch length - i - 1 {
		case 0:
			res[score[i]%100000] = "Gold Medal"
		case 1:
			res[score[i]%100000] = "Silver Medal"
		case 2:
			res[score[i]%100000] = "Bronze Medal"
		default:
			res[score[i]%100000] = strconv.Itoa(length - i)
		}
	}
	return res
}
