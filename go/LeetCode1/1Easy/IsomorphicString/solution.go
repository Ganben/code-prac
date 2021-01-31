package IsomorphicString

import "strconv"

func isIsomorphic(s string, t string) bool {
	sInt := numHelper(s)
	tInt := numHelper(t)
	for i, v := range sInt {
		if byte(tInt[i]) != byte(v) {
			return false
		}
	}
	return true
}

func numHelper(s string) string {
	m := map[int]int{}
	n := len(s)
	ret := ""
	count := 1
	for i := 0; i < n; i++ {
		c := int(s[i])
		if _, ok := m[c]; !ok {
			m[c] = count
			count++
		}
		ret += strconv.Itoa(m[c])
	}
	return ret
}
