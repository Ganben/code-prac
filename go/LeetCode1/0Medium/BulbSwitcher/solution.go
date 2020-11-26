package BulbSwitcher

func bulbSwitcher(n int) int {
	return sqrt(n)
}

func sqrt(x int) int {
	ans := 0
	for i := 1; i <= 1+x/2; i++ {
		if i*i <= x {
			ans++
		}
	}
	return ans
}
