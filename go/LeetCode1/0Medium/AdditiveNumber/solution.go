package AdditiveNumber

import "strconv"

func isAdditiveNumber(num string) bool {
	n := len(num)
	if n < 3 {
		return false
	}
	for i := 1; i < n; i++ {
		if i > 1 && num[0] == '0' {
			break
		}
		if lst, _ := strconv.Atoi(num[:i]); dfs(num, lst, i) {
			return true
		}
	}
	return false

}

func dfs(num string, lst int, idx int) bool {
	n := len(num)
	isAdditive := false
	for nxtStart := idx + 1; nxtStart < n; nxtStart++ {
		if nxtStart > idx+1 && num[idx] == '0' {
			break
		}
		cur, _ := strconv.Atoi(num[idx:nxtStart])
		sum := strconv.Itoa(cur + lst)
		dur := len(sum)
		nxtEnd := nxtStart + dur
		if nxtEnd > n {
			break
		}
		if sum == num[nxtStart:nxtEnd] {
			if nxtEnd == n {
				return true
			}
			isAdditive = dfs(num, cur, nxtStart)
		}
	}
	return isAdditive
}
