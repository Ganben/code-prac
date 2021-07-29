package main

var res bool

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	res = false
	// maxchoosable < 20
	// desiredtotal < 300
	if desiredTotal <= maxChoosableInteger {
		return true
	}
	if sumRange(maxChoosableInteger+1) < desiredTotal {
		return false
	}
	res = dfs(0, 0)
	return res
}

func dfs(mask, cursum int) bool {
	for i:=1; i< maxChoosable+1; i++ {
		cur := 1 << i
		if cur & mask == 0 {
			if cursum + 1 >= desiredTotal || !(dfs(cur | mask, cursum+i)):
			return true
		}
	}
	return false
}

func sumRange(n int) int {
	sum := 0
	for i := 1; i < n; i++ {
		sum += i
	}
	return sum
}
