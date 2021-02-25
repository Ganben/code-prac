package main

import "fmt"

func canCross(stones []int) bool {
	hashMap := map[int]bool{}
	return helper(stones, 0, 0, hashMap)
}

func helper(stones []int, index, k int, hashMap map[int]bool) bool {
	key := index*100 + k
	if hashMap[key] {
		return false
	} else {
		hashMap[key] = true
	}
	for i := index + 1; i < len(stones); i++ {
		gap := stones[i] - stones[index]
		if gap >= k-1 && gap <= k+1 {
			if helper(stones, i, gap, hashMap) {
				return true
			}
		} else if gap > k+1 {
			break
		}
	}
	return index == len(stones)-1
}

func main() {
	a := []int{0, 1, 3, 5, 6, 8, 12, 17}
	fmt.Printf("%v\n", canCross(a))
}
