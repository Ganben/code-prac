package main

import "fmt"

const (
	SMALLBIG = 0
	BIGSMALL = 1
	OTHER    = 2
)

func maxEnvelopes(envelopes [][]int) int {
	idArr := []int{}
	minEn := 0
	for i, v := range envelopes {
		if i == 0 {
			p := canEnvelop(envelopes[0], envelopes[1])
			if p == SMALLBIG {
				minEn = 0
				idArr = append(idArr, 0)
				idArr = append(idArr, 1)
			} else if p == BIGSMALL {
				minEn = 1
				idArr = append(idArr, 1)
				idArr = append(idArr, 0)
			}
		}
		// since 2
		if len(idArr) == 0 {

		}
	}
}

func canEnvelop(env1, env2 []int) int {
	if env1[0] < env2[0] && env1[1] < env2[1] {
		return SMALLBIG
	}
	if env1[0] < env2[0] && env1[1] < env2[1] {
		return BIGSMALL
	}
	return OTHER
}

func main() {
	enves := [][]int{
		{5, 4},
		{6, 4},
		{6, 7},
		{2, 3},
	}
	fmt.Printf("%v", maxEnvelopes(enves))
}
