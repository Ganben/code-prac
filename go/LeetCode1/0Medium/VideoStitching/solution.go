package VideoStitching

func videoStitching(clips [][]int, T int) int {
	maxn := make([]int, T)
	last, pre := 0, 0
	ans := 0
	for _, c := range clips {
		l, r := c[0], c[1]
		if l < T && r > maxn[l] {
			maxn[l] = r
		}
	}

	for i, v := range maxn {
		if v > last {
			last = v
		}
		if i == last {
			return -1
		}
		if i == pre {
			ans++
			pre = last
		}
	}
	return ans

}
