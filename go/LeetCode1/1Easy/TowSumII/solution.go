package TowSumII

func twoSum(numbers []int, target int) []int {
	a, b := 0, len(numbers)-1
	for a < b {
		if numbers[a]+numbers[b] == target {
			return []int{a, b}
		}
		d1 := numbers[a+1] + numbers[b] - target
		d2 := numbers[a] + numbers[b-1] - target
		if abs(d1) < abs(d2) {
			a++
		} else {
			b--
		}
	}
	return []int{-1, -1}
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -1 * x
}
