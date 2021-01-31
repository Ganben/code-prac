package main

import "fmt"

func restoreArray(adjacentPairs [][]int) []int {
	hmap := map[int][]int{}
	for _, el := range adjacentPairs {
		hmap[el[0]] = append(hmap[el[0]], el[1])
		hmap[el[1]] = append(hmap[el[1]], el[0])
	}
	// find two end
	// fistkey
	firstkey := 0
	for k, v := range hmap {
		if len(v) == 1 {
			firstkey = k
			break
		}
	}
	fmt.Printf("%v", hmap)
	ans := []int{firstkey}
	for i := 1; i < len(adjacentPairs)+1; i++ {
		next := hmap[firstkey]
		if len(next) == 1 {
			ans = append(ans, next[0])
		} else {
			if next[0] != ans[len(ans)-2] {
				ans = append(ans, next[0])
			} else {
				ans = append(ans, next[1])
			}
		}
		fmt.Printf("%v\n", ans)
		firstkey = ans[len(ans)-1]
		fmt.Printf("%v\n", firstkey)
	}
	return ans

}

func restoreArray2(adjacentPairs [][]int) []int {

	ans := []int{}
	vis := []bool{}
	for i := 0; i < len(adjacentPairs); i++ {
		vis = append(vis, false)
	}

	p1 := 1
	if len(adjacentPairs) == 0 {
		return ans
	}
	if len(adjacentPairs) == 1 {
		return adjacentPairs[0]
	}
	for p1 < len(adjacentPairs) {
		// find first two
		if p1 > len(adjacentPairs)-1 {
			p1 -= len(adjacentPairs)
		}
		if vis[p1] {
			p1++
			continue
		}
		//
		vis[0] = true // always skip 0 for starting point
		_, has := hasCommon(adjacentPairs[0], adjacentPairs[p1])
		if has == -1 {
			p1++
			continue
		} else if has == 0 {
			// left add
			// p0 = p1
			vis[p1] = true
			vis[0] = true
			if adjacentPairs[p1][0] != adjacentPairs[0][0] {
				ans = append([]int{adjacentPairs[p1][0]}, adjacentPairs[0]...)
			} else if adjacentPairs[p1][1] != adjacentPairs[0][0] {
				ans = append([]int{adjacentPairs[p1][1]}, adjacentPairs[0]...)
			}
			// next
			p1++
			continue
		} else if has == 1 {
			// left add
			vis[p1] = true
			vis[0] = true
			if adjacentPairs[p1][0] != adjacentPairs[0][1] {
				ans = append(adjacentPairs[0], adjacentPairs[p1][0])
			} else if adjacentPairs[p1][1] != adjacentPairs[0][1] {
				ans = append(adjacentPairs[0], adjacentPairs[p1][1])
			}
			// next
			p1++
			continue
		}
		// build first 3

	}

	// search
	p1 = 1
	flag := true
	for len(ans) < len(adjacentPairs)+1 {
		// build next, compare 2 end
		if flag {
			p1++
		} else {
			p1--
		}
		if p1 >= len(adjacentPairs)-1 {
			flag = false
		} else if p1 <= 1 {
			flag = true
		}

		if vis[p1] {
			continue
		}
		e1 := endCommon(ans[0], adjacentPairs[p1])
		e2 := endCommon(ans[len(ans)-1], adjacentPairs[p1])
		if e1 == -1 && e2 == -1 {
			continue
		} else if e1 == 0 {
			ans = append([]int{adjacentPairs[p1][1]}, ans...)
			vis[p1] = true

			continue
		} else if e1 == 1 {
			ans = append([]int{adjacentPairs[p1][0]}, ans...)
			vis[p1] = true
			continue
		} else if e2 == 0 {
			ans = append(ans, adjacentPairs[p1][1])
			vis[p1] = true
			continue
		} else if e2 == 1 {
			ans = append(ans, adjacentPairs[p1][0])
			vis[p1] = true
			continue
		} else {
			continue
		}
	}

	return ans
}

func endCommon(a1 int, ar2 []int) int {
	// which is common, -1 has no common
	if a1 == ar2[0] {
		return 0
	}
	if a1 == ar2[1] {
		return 1
	}
	return -1
}

func hasCommon(ar1, ar2 []int) (int, int) {
	//seq
	// 0 = left add
	// 1 = right add
	// compare
	if ar1[0] == ar2[0] {
		return ar1[0], 0
	}
	if ar1[0] == ar2[1] {
		return ar1[0], 0
	}
	if ar1[1] == ar2[0] {
		return ar1[1], 1
	}
	if ar1[1] == ar2[1] {
		return ar1[1], 1
	}
	return 0, -1
}

func main() {
	adjs := [][]int{
		{2, 1},
		{3, 4},
		{3, 2},
	}
	fmt.Printf("%v\n", restoreArray(adjs))
}
