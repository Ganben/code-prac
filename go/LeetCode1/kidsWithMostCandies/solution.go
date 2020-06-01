package solution

func kidsWithCandies(candies []int, extraCandies int) []bool {
	extra := make([]int, len(candies))
	result := make([]bool, len(candies))
	for i, n := range candies {
		extra[i] = n + extraCandies
	}
	for ii, _ := range candies {
		v := extra[ii]
		v1 := 0
		if ii > 0 {
			v1 = arrayMax(candies[:ii])
		}
		v2 := 0
		if ii+1 < len(candies) {
			v2 = arrayMax(candies[ii+1:])
		}

		if v > v1 && v > v2 {
			result[ii] = true
		}
	}
	return result

}

func arrayMax(ar []int) int {
	maxv := 0
	for _, n := range ar {
		if maxv < n {
			maxv = n
		}
	}
	return maxv
}
