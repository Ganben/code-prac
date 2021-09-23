package keyboardrow

func findWords(words []string) []string {
	ans := []string{}
	rows := []string{
		"qwertyuiop",
		"asdfghjkl",
		"zxcvbnm",
	}
	rowsMaps := []map[byte]bool{}
	for _, rs := range rows {
		mapT := map[byte]bool{}
		for _, ch := range rs {
			mapT[byte(ch)] = true
		}
		rowsMaps = append(rowsMaps, mapT)
	}
	for _, wd := range words {
		for _, mp := range rowsMaps {
			if search(mp, wd) {
				ans = append(ans, wd)
			}
		}
	}
	return ans
}

func search(mapH map[byte]bool, target string) bool {

	for _, ch := range target {
		if ch < 'a' {
			ch += 'a' - 'A'
		}
		if _, exist := mapH[byte(ch)]; !exist {
			return false
		}

	}
	return true
}
