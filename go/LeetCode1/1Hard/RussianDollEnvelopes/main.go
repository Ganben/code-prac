package main

import "fmt"

const (
	SMALLBIG     = 0
	BIGSMALL     = 1
	SAMEBUTBIG   = 2
	SAMEBUTSMALL = 3
	WSMALL       = 4
	WHCROSSED    = 5
)

func maxEnvelopes(envelopes [][]int) int {
	idArr := [][]int{}
	// minEn := 0
	// lis without remove sequence
	for i, v := range envelopes {
		if i == 0 {
			idArr = append(idArr, v)
			continue
		}
		//
		p := canEnvelop(idArr[len(idArr)-1], v)
		if p == SMALLBIG || p == SAMEBUTBIG || p == WSMALL {
			idArr = append(idArr, v)
			fmt.Printf("%v-1\n", idArr)
		} else {
			// bin search, greedy
			l := 0
			r := len(idArr) - 1
			pos := r
			// insert := true
			// replace := false
			for l <= r {
				mid := (l + r) / 2
				p1 := canEnvelop(idArr[mid], v)
				if p1 == BIGSMALL || p1 == SAMEBUTSMALL || p1 == WHCROSSED {
					// continue search
					pos = mid
					r = mid - 1
				} else {
					l = mid + 1
				}
			}
			// update idArr
			// if replace {
			// 	idArr[pos] = v
			// }
			// append, swap if needed

			head := [][]int{}
			for _, m := range idArr[:pos] {
				head = append(head, m)
			}
			tail := idArr[pos:]
			fmt.Printf("%v,\n", tail)
			head = append(head, v)
			fmt.Printf("%v,\n", tail)
			head = append(head, tail...)
			idArr = head

			fmt.Printf("%v-2\n", idArr)

		}
	}
	lis := []int{}
	// jump model
	for i, v := range idArr {
		// second time loop
		if i == 0 {
			lis = append(lis, v[1])
			continue
		}
		//
		if lis[len(lis)-1] < v[1] {
			lis = append(lis, v[1])
			fmt.Printf("%v\n", lis)
		} else {
			l := 0
			r := len(lis) - 1
			pos := r
			for l <= r {
				mid := (l + r) / 2
				if lis[mid] >= v[1] {
					pos = mid
					r = mid - 1
				} else {
					l = mid + 1
				}
			}
			lis[pos] = v[1]
		}
		// case 4

	}
	return len(lis)
}

func canEnvelop(env1, env2 []int) int {
	if env1[0] < env2[0] && env1[1] <= env2[1] {
		return SMALLBIG
	}
	//
	if env1[0] > env2[0] && env1[1] >= env2[1] {
		return BIGSMALL
	}
	//
	if env1[0] == env2[0] {
		if env1[1] > env2[1] {
			return SAMEBUTBIG
		} else {
			return SAMEBUTSMALL
		}
	}
	// other, crossed
	if env1[0] > env2[0] {
		return WHCROSSED
	}
	return WSMALL
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
