package main

import (
	"fmt"
	"sort"
)

func freuencySort(s string) string {
	hm := map[rune]int{}
	// hm2 := map[int]rune{}
	for _, c := range s {
		_, exist := hm[c]
		if exist {
			hm[c]++
		} else {
			hm[c] = 1
		}
		// v := hm[c]
		// hm2[v] =
	}
	//
	// nums := make([]int, len(hm))
	// // chars := make([]byte, len(hm))
	// idx := 0
	// for k, v := range hm {
	// 	nums[idx] = v
	// 	// chars[idx] = byte(k)
	// 	idx++
	// }
	// //
	// sort.Ints(nums)
	// for _, v := range nums {

	// }
	type sChar struct {
		Char byte
		Val  int
	}
	numbered := []sChar{}
	for k, v := range hm {
		numbered = append(numbered, sChar{byte(k), v})
	}
	sort.Slice(numbered, func(i, j int) bool {
		return numbered[i].Val > numbered[j].Val
	})
	ans := ""
	for _, em := range numbered {
		tmp := ""
		for i := 0; i < em.Val; i++ {
			tmp += string(em.Char)
		}
		ans += tmp
	}
	return ans
}

func main() {
	i := "tree"
	fmt.Printf("%v\n", freuencySort(i))
}
