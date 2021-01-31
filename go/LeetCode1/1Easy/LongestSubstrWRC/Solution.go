package Solution

// O(n) time O(1) space
func lengthOfLongestSubstring(s string) int {
	var chPosition [256]int //
	maxLength, substrLen, lastRepeatPos := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if pos := chPosition[s[i]]; pos > 0 {
			//
			maxLength = Max(substrLen, maxLength)
			//
			chPosition[s[i]] = i + 1
			//
			lastRepeatPos = Max(pos, lastRepeatPos)

			// 
			substrLen = i + 1 - lastRepeatPos
		} else {
			substrLen += 1
			chPosition[s[i]] = i + 1
		}
	}
	return Max(maxLength, substrLen)
}


func lengthOfLongestSubstring2(s string) int {
	ans := 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if allUnique(s, i, j) {
				ans = Max(ans, j-i)
			}
		}
	}
	return ans
}

func allUnique(s string, start int, end int) bool {
	sMap := make(map[string]int)
	for i := start; i < end; i++ {
		if sMap[string(s[i])] > 0 {
			return false
		}
		sMap[string(s[i])]++
	}
	return true
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
