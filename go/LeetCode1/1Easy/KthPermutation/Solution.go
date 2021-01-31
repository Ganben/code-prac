package solution

var ans string
var visit []bool

var count int

func getPermutation(n, k int) string {
	ans = ""
	count = 0
	visit = make([]bool, n)
	// start pre search
	search(n, k, []int{})
	if count >= k {
		return ans
	}
	return ""
}

func search(n, k int, curr []int) {
	// m means m th 1 level sub division for n combs
	if len(curr) == n {
		count++
		s := ""
		// fmt.Printf("%v\n", curr)
		for _, c := range curr {
			// s += strconv.Itoa(c)
			s += string(48 + c)
			// fmt.Printf("add %v \n", s)
		}
		ans = s
		return
	}

	p := len(curr)
	pan := combinations(n - p - 1)
	rest := k - count

	if rest <= pan {
		// inside this pan
		for i, x := range visit {
			if !x {
				curr = append(curr, i+1)
				visit[i] = true
				// fmt.Printf("deeper at %v\n", curr)
				search(n, k, curr)
			}
		}
	} else {
		// move to nth of sub div
		m := int(rest / pan)
		// deduct the pan * nth div
		if rest%pan > 0 {
			count += pan * m
		} else {
			m--
			count += pan * m
		}

		//

		nth := 0
		for i, x := range visit {
			if !x {
				nth++
			}
			if nth == m+1 {
				curr = append(curr, i+1)
				visit[i] = true
				// fmt.Printf("deeper at %v, for %v\n", curr, nth)
				search(n, k, curr)
				break
			}
		}

	}

}

func combinations(n int) int {
	res := 1
	for i := 1; i < n+1; i++ {
		res = res * i
	}
	return res
}
