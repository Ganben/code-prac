package Combinations

var ans [][]int
var visited []bool

func combine(n, k int) [][]int {
	ans = [][]int{}
	curr := []int{}
	visited = make([]bool, n)
	if k > 0 && n >= 1 {
		walker(n, curr, k)
	}
	return ans
}

func walker(n int, curr []int, l int) {
	if len(curr) == l {
		tmp := make([]int, l)
		copy(tmp, curr)
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
