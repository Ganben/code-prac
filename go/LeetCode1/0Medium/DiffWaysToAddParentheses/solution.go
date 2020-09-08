package DiffWaysToAddParentheses

import "strconv"

func diffWaysToCompute(input string) []int {
	if isDigit(input) {
		tmp, _ := strconv.Atoi(input)
		return []int{tmp}
	}
	var res []int
	for index, c := range input {
		tmpchar := string(c)
		if tmpchar == "+" || tmpchar == "-" || tmpchar == "*" {
			left := diffWaysToCompute(input[:index])
			right := diffWaysToCompute(input[index+1:])

			for _, leftNum := range left {
				for _, rightNum := range right {
					var addN int
					if tmpchar == "+" {
						addN = leftNum + rightNum
					} else if tmpchar == "-" {
						addN = leftNum - rightNum
					} else {
						addN = leftNum * rightNum
					}
					res = append(res, addN)
				}
			}
		}
	}
	return res

}

func isDigit(input string) bool {
	for _, v := range input {
		b := int(v - '0')
		if b > 9 || b < 0 {
			return false
		}
	}
	return true
}
