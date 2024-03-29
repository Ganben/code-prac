package longestwordindict

func findLongestWord(s string, dictionary []string) string {
	ans := ""
	for _, t := range dictionary {
		i := 0
		for j := range s {
			if s[j] == t[i] {
				i++
				if i == len(t) {
					if len(t) > len(ans) || len(t) == len(ans) && t < ans {
						ans = t
					}
					break
				}
			}
		}
	}
	return ans
}
