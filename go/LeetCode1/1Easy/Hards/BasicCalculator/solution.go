package BasicCalculator

import "container/list"

// var stack list
// var stack list

func calculate(s string) int {
	stack := list.New()
	operand := 0
	result := 0
	sign := 1

	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch-'0' >= 0 && ch-'0' < 10 {
			operand = 10*operand + int(ch-'0')
		} else if ch == '+' {
			result += sign * operand
			sign = 1
			operand = 0
		} else if ch == '-' {
			result += sign * operand
			sign = -1
			operand = 0
		} else if ch == '(' {
			stack.PushBack(result)
			stack.PushBack(sign)
			sign = 1
			result = 0
		} else if ch == ')' {
			result += sign * operand
			stackSign := stack.Back()
			result *= stackSign.Value.(int)
			stack.Remove(stackSign)
			stackResult := stack.Back()
			result += stackResult.Value.(int)
			stack.Remove(stackResult)
			operand = 0
		}
	}
	return result + (sign * operand)
}
