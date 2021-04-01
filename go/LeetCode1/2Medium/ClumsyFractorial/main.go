package main

func clumsy(N int) int {
	//
	ans := 1
	ansArr := []int{}
	if N <= 3 {
		for i := 0; i < N; i++ {
			if i == 0 {
				ans = N
			} else if i == 1 {
				ans *= N - i
			} else if i == 2 {
				ans /= N - i
			}
		}
		return ans
	}
	ans = 0
	for i := 2; i < N; i++ {
		if i == 2 {
			ansArr = append(ansArr, N*(N-1)/(N-2))
			continue
		}
		if i%4 == 0 && N-i > 2 {
			continue
		} else if i%4 == 1 && N-i > 1 {
			continue
		} else if i%4 == 2 && N-i >= 1 {
			ansArr = append(ansArr, -1*(N-i+2)*(N-i+1)/(N-i))
			continue
		} else if i%4 == 3 {
			ansArr = append(ansArr, N-i)
			continue
		}
		//final loop
		if i%4 == 0 && N-i == 1 {
			ansArr = append(ansArr, -1*(N-i))
		}
		if i%4 == 1 && N-i == 1 {
			ansArr = append(ansArr, -1*(N-i+1)/(N-i))
		}

	}
	for _, v := range ansArr {
		ans += v
	}
	return ans
}
