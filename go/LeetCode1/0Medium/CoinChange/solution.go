package CoinChange

import "sort"

func coinChange(coins []int, amount int) int {
	// coinCount := make([]int, len(coins))
	sort.Ints(coins)
	p1, rest := len(coins)-1, amount
	ans := []int{}
	for p1 >= 0 {
		deduct := rest / coins[p1]
		rest -= deduct * coins[p1]
		ans = append(ans, deduct)
		if rest == 0 {
			break
		}
		p1--
	}
	//back search required
	if rest != 0 {
		return -1
	}
	res := 0
	for _, v := range ans {
		res += v
	}
	return res
}
