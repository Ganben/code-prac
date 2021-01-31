package WordBreakII

var hashMap map[int][]string

func wordBreak(s string, wordDict []string) []string {
	hashMap = map[int][]string{}
	return helper(s, wordDict, 0)
}

func helper(s string, wordDict []string, start int) []string {
	if _, ok := hashMap[start]; ok {
		return hashMap[start]
	}
	res := []string{}
	if start == len(s) {
		res = append(res, "")
	}
	for end := start + 1; end <= len(s); end++ {
		if listHas(wordDict, s[start:end]) {
			list := helper(s, wordDict, end)
			for _, l := range list {
				res = append(res, s[start:end]+makeL(l)+l)
			}
		}

	}
	hashMap[start] = res
	return res
}

func listHas(wordDict []string, substr string) bool {
	for _, v := range wordDict {
		if v == substr {
			return true
		}
	}
	return false
}

func makeL(l string) string {
	if l == "" {
		return ""
	}
	return " "
}
