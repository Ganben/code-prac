package CombinationSumiii

var result [][]int
var nN int
var kN int

func combinationSum3(k int, n int) [][]int {
	result = [][]int{}
	nN = n
	kN = k

	dfs(0, 0, []int{})
	return result
}

func dfs(numN, sumN int, tmp []int) {
	if len(tmp) > kN || sumN > nN {
		return
	}
	if len(tmp) == kN && sumN == nN {
		t := make([]int, len(tmp))
		copy(t, tmp)
		result = append(result, t)
		return
	}
	for i := numN + 1; i < 10; i++ {
		if sumN+i > nN {
			break
		}
		dfs(i, sumN+i, append(tmp, i))
	}

}
