package BestTimeBuyStock

func maxProfit(prices []int) int {
	minprice := sum(prices)
	maxProfit := 0
	for _, v := range prices {
		maxProfit = max(v-minprice, maxProfit)
		minprice = min(v, minprice)
	}
	return maxProfit
}

func sum(ar []int) int {
	ans := 0
	for _, v := range ar {
		ans += v
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
