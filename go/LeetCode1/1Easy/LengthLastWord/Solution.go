package solution

func lengthOfLastWord(s string) int {
	n := 0
	flag := false
	for _, char := range s {
		if char == ' ' {
			flag = true
			continue
		}
		if flag {
			n = 0
			flag = false
		}
		n++
	}
	return n
}
