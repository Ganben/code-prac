package SuperUglyNumber

import "sort"

func nthSuperUglyNumber(n int, primes []int) int {
	mutation := append(primes, 1)
	multi := make([]int, len(primes))
	for i, _ := range multi {
		multi[i] = 1
	}
	cur := 0
	for cur < primes[len(primes)-1] {

		for i, _ := range multi {

			d := power(primes[i], multi[i])
			if i == 0 {
				cur = d
			}
			multi[i] += 1
			if d < primes[0]*primes[len(primes)-1] {
				mutation = append(mutation, d)
			}
		}

	}
	sort.Ints(mutation)
	mmap := map[int]int{}
	for _, v := range mutation {
		mmap[v] = 1
	}
	lenRow := 1
	for lenRow < n {
		row := map[int]int{}
		for _, v1 := range mutation[1:] {
			for _, v2 := range primes {
				row[v1*v2] = 1
			}
		}
		for k, _ := range row {
			mmap[k] = 1
		}
		lenRow = len(row)
	}

	keys := []int{}
	for k, _ := range mmap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	return keys[n-1]
}

func power(x, y int) int {
	ans := 1
	for i := 0; i < y; i++ {
		ans *= x
	}
	return ans
}
