package base7

import "strconv"

func convertToBase7(num int) string {
	digits := []int{}
	sign := ""
	curr := num
	if num < 0 {
		curr = -1 * num
		sign = "-"
	}

	for curr >= 7 {
		digits = append(digits, curr%7)
		curr = curr / 7
	}
	digits = append(digits, curr)
	ans := sign
	for i := 0; i < len(digits); i++ {
		ans += strconv.Itoa(digits[len(digits)-1-i])
	}

	return ans
}
