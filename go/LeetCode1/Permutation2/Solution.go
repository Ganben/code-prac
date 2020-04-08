package Solution

var ans [][]int
var visit []int
var numset [][]int

func permuteUnique(nums []int) [][]int {
	numset = [][]int{}
	for _, num := range nums {
		putSet(num)
	}
	l := len(nums)
	ans = [][]int{}
	visit = make([]int, len(numset))
	backtrack(l, []int{})
	return ans
}

func backtrack(l int, curr []int) {
	if len(curr) == l {
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		ans = append(ans, tmp)
		return
	}

	for idx, c := range numset {
		if visit[idx] < c[1] {
			curr = append(curr, c[0])
			visit[idx]++
			backtrack(l, curr)
			visit[idx]--
			curr = curr[:len(curr)-1]
		}

	}
}

func putSet(i int) {

	for _, c := range numset {
		if i == c[0] {
			c[1]++
			return
		}
	}
	numset = append(numset, []int{i, 1})
}
