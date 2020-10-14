package FindCommonChar

import "math"

func commonChars(A []string) []string {
	ans := []string{}
	for _, e := range A[0] {
		f := true
		for i := 1; i < len(A); i++ {
			if test(e, A[i]) {
				f = false
				break
			}
		}
		if f {
			ans = append(ans, string(e))
		}

	}
	return ans
}

func test(e rune, a string) bool {
	for _, c := range a {
		if e == c {
			return false
		}
	}
	return true
}

func res(a []string) []string {
	ans := []string{}
	minFreq := [26]int{}
	for i := range minFreq {
		minFreq[i] = math.MaxInt64
	}
	for _, word := range a {
		freq := [26]int{}
		for _, b := range word {
			freq[b-'a']++
		}
		for i, f := range freq[:] {
			if f < minFreq[i] {
				minFreq[i] = f
			}
		}
	}
	for i := byte(0); i < 26; i++ {
		for j := 0; j < minFreq[i]; j++ {
			ans = append(ans, string('a'+i))
		}
	}
	return ans

}
