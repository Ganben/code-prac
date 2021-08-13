package main

func findComplement(num int) int {
	ans := 0
	index := 0
	for num != 0 {
		ans = ans + (num&1^1)<<index
		num >>= 1
		index++
	}
	return ans
}
