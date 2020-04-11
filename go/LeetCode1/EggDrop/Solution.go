package Solution

var dp = [][]int{}

func superEggDrop(K, N int) int {
	dp = [][]int{}
	// return binsearch(K, N)
	// init dp array

	for i := 0; i < K+1; i++ {
		row0 := make([]int, N+1)
		dp = append(dp, row0)
	}

	for i := 1; i < K+1; i++ {
		for j := 1; j < N+1; j++ {
			// fmt.Printf("fill %v,%v\n", i, j)
			dpsearch(i, j)
		}
	}
	// fmt.Printf("%v", dp)
	// fmt.Printf("==============%v, %v return %v============\n", K, N, dp[K][N])
	return dp[K][N]
}

func dpsearch(K, N int) {

	if K == 1 {
		dp[K][N] = N
	} else {
		if N == 1 {
			dp[K][N] = 1
		} else {
			tmp := dp[K-1][N]
			// tmax := 0
			lo, hi := 1, N
			mid, b1, b2 := 0, 0, 0

			for lo <= hi {
				if K > 1 {
					// dp[0][i] = 0
					mid = int((lo + hi) / 2)
					b1 = dp[K-1][mid-1]
					b2 = dp[K][N-mid]
					if b1 > b2 {
						hi = mid - 1
						tmp = min(tmp, b1+1)
					} else {
						lo = mid + 1
						tmp = min(tmp, b2+1)
					}

					// tmax = 1 + max(dp[K-1][i-1], dp[K][N-i])
					// if i == 1 {
					// 	tmp = tmax
					// } else {
					// 	if tmp > tmax && tmax > 0 {

					// 		tmp = tmax
					// 	}
					// }
				}
			}
			// fmt.Printf("bettered %v, %v with %v \n", K, N, tmp)
			dp[K][N] = tmp
		}
	}

	//

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func binsearch(K, N int) int {
	// binsearch is limited by the K/N ratio, if very diff, the strategy
	// will not work (unresolved)
	if K > N {
		K = N
	}

	if K == 1 {
		if N == 1 {
			return 1
		}
		return N
	}

	// recursive, make worse k - n ratio (in binary search)
	// c := 0
	if N%2 == 0 {
		N = int(N / 2)
		return 1 + binsearch(K, N)
	} else {
		N = int(N / 2)
		p1 := 1 + binsearch(K-1, N)
		p2 := 1 + binsearch(K, N+1)
		if p1 > p2 {
			return p2
		}
		return p1
	}
}
