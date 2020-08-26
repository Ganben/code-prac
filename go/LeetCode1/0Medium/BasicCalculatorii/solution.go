package BasicCalculatorii

import "container/list"

func calculate(s string) int {
	stack := list.New()
	operand := 0
	symb := '_'
	// lastSymb := '_'
	stacked := false
	result := []int{}
	sign := 1

	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch == ' ' {
			continue
		}

		if ch-'0' >= 0 && ch-'0' < 10 {
			operand = 10*operand + int(ch-'0')
			stacked = false
		} else {
			if symb != '_' {
				if ch == '-' && operand == 0 {
					sign = -1
					continue
				} else {
					if symb == '*' {
						t := result[len(result)-1]
						result = result[:len(result)-1]
						result = append(result, sign*operand*t)
						operand = 0
						sign = 1
						symb = '_'
						stacked = true

					} else if symb == '/' {
						t := result[len(result)-1]
						result = result[:len(result)-1]
						result = append(result, sign*int(t/operand))
						operand = 0
						sign = 1
						symb = '_'
						stacked = true

					}

					if ch == '*' {
						symb = '*'
						continue
					}

					if ch == '/' {
						symb = '/'
						continue
					}
				}

			}

			if ch == '+' {
				result = append(result, sign*operand)
				sign = 1
				operand = 0
			} else if ch == '-' {
				result = append(result, sign*operand)
				sign = -1
				operand = 0
			} else if ch == '(' {
				// stack.PushBack(operand)
				stack.PushBack(sign)
				sign = 1
				// operand = 0
				// result = 0
			} else if ch == ')' {

				stackSign := stack.Back()
				sx := stackSign.Value.(int)
				result = append(result, sx*operand)
				stack.Remove(stackSign)

				operand = 0
			} else if ch == '*' {
				// t := result[len(result)-1]
				// result = result[:len(result)-1]
				if !stacked {
					result = append(result, sign*operand)
					sign = 1
					symb = '*'
					operand = 0

				}

			} else if ch == '/' {
				if !stacked {

					result = append(result, sign*operand)
					sign = 1
					symb = '/'
					operand = 0

				}
			}

		}
	}
	ans := 0
	for _, v := range result {
		ans += v
	}
	if symb == '_' {
		return ans + (sign * operand)
	}

	t := result[len(result)-1]

	if symb == '*' {
		ans -= t
		ans += sign * t * operand
	} else if symb == '/' {
		ans -= t
		ans += sign * int(t/operand)
	}

	return ans

}
