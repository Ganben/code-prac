package zumagame

func findMinStep(board string, hand string) int {
	handMap := make(map[byte]int)
	handLen := len(hand)
	for i := 0; i < handLen; i++ {
		handMap[hand[i]]++
	}
	return dfs(board, handMap)
}

func dfs(board string, hand map[byte]int) int {
	n := len(board)
	if board == "" {
		return 0
	}
	for i := 0; i < n; i++ {
		j := i
		for i+1 < n && board[i+1] == board[j] {
			i++
		}
		if i-j+1 >= 3 && removeDup(board, i) == "" {
			return 0
		}
	}
	cnt := -1
	for i := 0; i < n; i++ {
		j := i
		for i+1 < n && board[i+1] == board[j] {
			i++
		}
		need := 3 - (i - j + 1)
		if need > 0 && hand[board[i]] >= need {
			hand[board[i]] -= need
			if v := dfs(board[:j]+board[i+1:], hand); v >= 0 && (cnt < 0 || v < cnt) {
				cnt = v + need
			}
			hand[board[i]] += need
		}
	}
	return cnt

}

func removeDup(s string, i0 int) string {
	for len(s) > 0 {
		ln := len(s)
		i, j := i0, i0
		for i-1 >= 0 && s[i-1] == s[i0] {
			i--
		}
		for j+1 < ln && s[j+1] == s[i0] {
			j++
		}
		if j-i+1 < 3 {
			return s
		}
		s = s[:i] + s[j+1:]
		i0 = max(i-1, 0)
	}
	return s
}

func max(x, y int) int {
	if y > x {
		return y
	}
	return x
}

func min(x, y int) int {
	if y < x {
		return y
	}
	return x
}
