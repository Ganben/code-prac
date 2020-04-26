package solution

func plusOne(digits []int) []int {
	ans := []int{}
	for i, n := range digits {
		if i == len(digits)-1 {
			if n+1 == 10 {
				digits[i] = 0
				if i-1 >= 0 {
					ans = plusOne(digits[:i])
					ans = append(ans, 0)
					break
				} else {
					ans = append(ans, 1)
					ans = append(ans, 0)
					break
				}
			}
			digits[i] = n + 1
			ans = append(ans, digits...)
		}
	}
	return ans
}
