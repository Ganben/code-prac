package MaximumProductOfWordLengths

func maxProduct(words []string) int {
	ans := 0
	n := len(words)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if ifCommon(words[i], words[j]) {
				if len(words[i])*len(words[j]) > ans {
					ans = len(words[i]) * len(words[j])
				}
			}
		}
	}
	return ans
}

func ifCommon(w1, w2 string) bool {
	vmap1 := map[rune]int{}
	for _, c := range w1 {
		vmap1[c] = 1
	}
	for _, c := range w2 {
		_, exist := vmap1[c]
		if exist {
			return false
		}
	}
	return true
}
