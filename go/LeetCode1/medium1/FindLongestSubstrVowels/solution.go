package FindLongestSubstrVowels

func findTheLongestSubstring(s string) int {
	vcount := make([]int, 5)
	for _, c := range s {
		switch c {
		case 'a':
			vcount[0]++
		case 'e':
			vcount[1]++
		case 'i':
			vcount[2]++
		case 'o':
			vcount[3]++
		case 'u':
			vcount[4]++
		}
	}

}

func isEven(vcount []int) bool {
	f := false
	for _, c := range vcount {
		if c != 0 {
			f = true
		}
		if c%2 != 0 {
			return false
		}
	}
	if f {
		return true
	}
}
