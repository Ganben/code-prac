package solution

func dailyTemperatures(T []int) []int {
	l := len(T)
	ans := make([]int, l)
	stack := []int{}
	for i := 0; i < l; i++ {
		temperature := T[i]
		for len(stack) > 0 && temperature > T[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}
	return ans
}
