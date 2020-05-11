package Subsets

var ans [][]int
var visited []bool
var numSets []int

func subsets(nums []int) [][]int {
	ans = [][]int{}
	curr := []int{}
	n := len(nums)
	visited = make([]bool, n)
	numSets = make([]int, n)
	for i, c := range nums {
		numSets[i] = c
	}

	for k := 0; k <= n; k++ {
		walker(n, curr, k)
	}

	return ans
}

func walker(n int, curr []int, l int) {
	if len(curr) == l {
		tmp := make([]int, l)
		for i, c := range curr {
			tmp[i] = numSets[c-1]
		}
		ans = append(ans, tmp)
		return
	}
	ps := 1
	if len(curr) == 0 {
		ps = 1
	} else {
		ps = curr[len(curr)-1]
	}

	for idx := ps; idx <= n; idx++ {
		if !visited[idx-1] {
			curr = append(curr, idx)
			visited[idx-1] = true
			walker(n, curr, l)
			visited[idx-1] = false
			curr = curr[:len(curr)-1]
		}
	}

}
