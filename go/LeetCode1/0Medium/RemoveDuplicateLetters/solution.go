package RemoveDuplicateLetters

import "sort"

type Cposition struct {
	Ch  string
	Pos int
}

func removeDuplicateLetters(s string) string {
	dictNum := map[int]int{}
	idArray := [][]int{}
	for i := 0; i < 26; i++ {
		row := make([]int, len(s))
		idArray = append(idArray, row)
	}

	for i, c := range s {
		// val, exist := dictNum[c]
		idArray[c-'a'][i] = 1 // presense map
	}
	ans := ""
	// seqs := []int{}
	for i, row := range idArray {
		x := []int{}
		for j, val := range row {
			if val == 1 {
				x = append(x, j)
			}
		}
		// decide position
		if len(x) == 0 {
			continue
		}
		if len(x) == 1 {
			// only choice
			dictNum[x[0]] = 97 + i
			continue
		}
		if len(x) > 1 {
			curP := x[0]
			for key, _ := range dictNum {
				for _, p := range x {
					if p > key {
						curP = p

					}
				}

			}
			dictNum[curP] = 97 + i
		}
	}
	sortArr := []Cposition{}
	for key, val := range dictNum {
		c := Cposition{
			Ch:  string(byte(val)),
			Pos: key,
		}
		sortArr = append(sortArr, c)
	}
	sort.Slice(sortArr, func(i, j int) bool {
		return sortArr[i].Pos < sortArr[j].Pos
	})
	for _, cp := range sortArr {
		ans += cp.Ch
	}
	return ans
}
