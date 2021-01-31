package ExprAddOperators

import (
	"strconv"
	"strings"
)

var (
	g_target int
	g_num    string
	g_result []string
)

const (
	add byte = '+'
	sub byte = '-'
	mul byte = '*'
)

func addOperators(num string, target int) []string {
	g_num = num
	g_target = target
	g_result = make([]string, 0)
	backtrack(0, 0, add, new([]string))
	return g_result
}

func backtrack(sum, i int, op byte, expr *[]string) {
	for j := i + 1; j <= len(g_num); j++ {
		if j > i+1 && g_num[i] == '0' {
			continue
		}
		*expr = append(*expr, string(op), g_num[i:j])
		if j == len(g_num) {
			if calc(expr) == g_target {
				g_result = append(g_result, strings.Join((*expr)[1:], ""))
			}
			*expr = (*expr)[:len(*expr)-2]
			return
		}
		backtrack(sum, j, add, expr)
		backtrack(sum, j, sub, expr)
		backtrack(sum, j, mul, expr)
		*expr = (*expr)[:len(*expr)-2]

	}

}

func calc(expr *[]string) int {
	stack := make([]int, 0)
	for i := 0; i < len(*expr); i += 2 {
		x, _ := strconv.ParseInt((*expr)[i+1], 10, 32)
		switch (*expr)[i] {
		case "+":
			stack = append(stack, int(x))
		case "-":
			stack = append(stack, -int(x))
		case "*":
			p := &stack[len(stack)-1]
			*p = *p * int(x)
		}
	}
	ans := 0
	for _, x := range stack {
		ans += x
	}
	return ans
}
