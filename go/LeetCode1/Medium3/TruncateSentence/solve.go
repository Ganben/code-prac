package truncatesentence

func truncateSentence(s string, k int) string {
	idx := 0
	for i, char := range s {
		if char == ' ' {
			idx++
			if idx == k {
				return s[:i]
			}
		}
	}
	return s
}
