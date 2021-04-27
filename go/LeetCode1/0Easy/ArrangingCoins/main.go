package main

func arrangeCoins(n int) int {
	remains := n
	ans := 0
	for i := 1; i < n+1; i++ {
		remains -= i
		ans++
		if remains < i+1 {
			break
		}

	}
	return ans
}
