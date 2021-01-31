package CountPrimes

import "math"

func countPrimes(n int) int {
	isPrime := make([]bool, n)
	for i, _ := range isPrime {
		isPrime[i] = true
	}
	for i := 2; i*i < n; i++ {
		if isPrime[i] {
			for j := i * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	count := 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
		}
	}
	return count
}

func sqrt(n int) int {
	return int(math.Sqrt(float64(n)))
}
