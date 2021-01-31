package ContestW

func minCost(s string, cost []int) int {
	if len(s) == 1 {
		return 0
	}
	//repeat starts and end ids
	starts := -1
	// ends := -1
	ans := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			if starts == -1 {
				starts = i
			}

		} else {
			if starts != -1 {

				ans += min(cost[starts : i+1])
				starts = -1
			}
		}
		// end of end
		if i == len(s)-2 && starts != -1 {
			ans += min(cost[starts:])
		}
	}
	return ans

}

func min(iarr []int) int {
	ans := 0
	mx := 0
	for _, v := range iarr {
		if v > mx {
			mx = v
		}
		ans += v
	}
	return ans - mx
}
