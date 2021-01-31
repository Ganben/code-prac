package Solution

var m = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	
	sum := m[string(s[len(s)-1])]
	// big end to small end
	for i := len(s) - 2; i>= 0; i-- {
		if m[string(s[i])] < m[string(s[i+1])] {
			sum -= m[string(s[i])]
		} else {
			sum += m[string(s[i])]
		}
	}
	return sum
}

func romanToInt2(s string) int {
	ans := 0
	// small end to big end
	for i, _ := range s {
		ans += m[string(s[i])]
		if i > 0 && m[string(s[i])] > m[string(s[i-1])] {
			ans -= 2 * m[string(s[i-1])]
		}
	}
	return ans
}