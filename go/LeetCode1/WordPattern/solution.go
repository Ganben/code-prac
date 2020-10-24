package WordPattern

import "strings"

func wordPattern(pattern string, s string) bool {
	// build dict
	pMap := map[byte]string{}
	// split string
	sArr := strings.Split(s, " ")
	if len(pattern) != len(sArr) {
		return false
	}
	for i := 0; i < len(pattern); i++ {
		p := byte(pattern[i])
		if _, found := pMap[p]; !found {
			pMap[p] = sArr[i]
		} else {
			if pMap[p] != sArr[i] {
				return false
			}

		}
		for k, v := range pMap {
			if v == sArr[i] && k != p {
				return false
			}
		}
	}
	return true

}
