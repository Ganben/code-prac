package main

func findMaximumXOR(nums []int) int {
	maxNum := nums[0]
	for _, i := range nums {
		maxNum = max(maxNum, i)
	}
	//length of maxNum in binary form(not null lenth)
	L := 0
	f := maxNum
	for f != 0 {
		L++
		f = f >> 1
	}
	//
	maxXor, currXor := 0, 0
	prefixes := map[int]int{}
	for i := L - 1; i > -1; i-- {
		maxXor <<= 1
		currXor = maxXor | 1
		//clear prefixes
		for _, num := range nums {
			prefixes[num>>i] = 1
		}
		//
		for k, _ := range prefixes {
			if prefixes[currXor^k] == 1 {
				maxXor = currXor
				break
			}
		}
	}
	return maxXor
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {

}
