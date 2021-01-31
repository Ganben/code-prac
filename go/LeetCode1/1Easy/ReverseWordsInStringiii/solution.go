package ReverseWordsInStringiii

func reverseWords(s string) string {
	i1, i2 := 0, 0
	arr := []string{}
	// split
	for i, v := range s {
		if v == ' ' {
			i1 = i2 + 1
			i2 = i
			if i2-i1 > 0 {
				arr = append(arr, s[i1:i2])
			}
			continue
		}
		if i == len(s)-1 {
			arr = append(arr, s[i2+1:])
		}
	}
	// reverse
	for i := 0; i < len(arr); i++ {
		arr[i] = revWord(arr[i])
	}

	// resemble
	result := ""
	for i, v := range arr {
		result += v
		if i < len(arr)-1 {
			result += " "
		}
	}
	return result
}

func revWord(s string) string {
	n := len(s)
	res := ""
	for i := 0; i < n; i++ {
		res += string(s[n-1-i])
	}
	return res
}
