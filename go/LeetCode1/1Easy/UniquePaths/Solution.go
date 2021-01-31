package solution

import "fmt"

func uniquePathsB(m, n int) int {
	// bfs/dfs ok
	result := 0
	if m > 1 {
		result += uniquePathsB(m-1, n)
	} else {
		return 1
	}

	if n > 1 {
		result += uniquePathsB(m, n-1)
	} else {
		return 1
	}
	return result
}

func uniquePaths(m, n int) int {

	return combination(m+n-2, m-1)
}

func combination(m, n int) int {
	var result int64
	var resultStack []int64
	resultStack = []int64{}
	result = 1

	if m == 0 {
		return 1
	}
	if m == 1 {
		return 1
	}
	if n == 0 {
		return 1
	}

	if n == 1 {
		return m
	}

	for i := m; i >= n+1; i-- {
		if result*int64(i) > 2147483647 {
			resultStack = append(resultStack, result)
			result = int64(i)
		} else {
			result = result * int64(i)
		}

		// fmt.Printf("big int %v where %v, stacked %v\n", result, i, resultStack)
	}

	if m-n == 0 {
		return int(result)
	}
	if len(resultStack) > 0 {
		resultStack = append(resultStack, result)
		dividStack := []int{}
		for i := 1; i <= m-n; i++ {
			flag := false
		Loop1:
			for j, bigNum := range resultStack {
				if bigNum%int64(i) == 0 {
					resultStack[j] = resultStack[j] / int64(i)
					flag = true
					break Loop1
				}
			}
			if !flag {
				fmt.Printf("no divid, stack %v to %v\n", i, dividStack)
				dividStack = append(dividStack, i)
			}
		}
		result = mergeStacks(resultStack, dividStack)

	} else {
		for i := m - n; i > 0; i-- {

			result = result / int64(i)
		}
	}

	return int(result)
}

func mergeStacks(resultStack []int64, dividStack []int) int64 {
	// merge first
	// for i, n := range resultStack {
	// 	if i+1 < len(resultStack) {
	// 		if n*resultStack[i+1] > n {
	// 			resultStack[i] = n * resultStack[i+1]
	// 			resultStack[i+1] = 1
	// 		}
	// 	}
	// }
	// search divid

	if len(dividStack) > 0 {
		for testDivide(dividStack) {
			for i, n := range resultStack {
				if n == 1 {
					continue
				}
				for j, d := range dividStack {
					if d != 1 && n%int64(d) == 0 {
						resultStack[i] = resultStack[i] / int64(d)
						dividStack[j] = 1
					} else {
						// fmt.Printf("failed div %v/%v=%v\n", n, d, n%int64(d))
					}
				}
			}
			// factors split
			for i := 0; i < len(resultStack); i++ {
				if i+1 < len(resultStack) {
					if resultStack[i+1]*resultStack[i] < 2147483647*2147483647 {

						resultStack[i] = resultStack[i+1] * resultStack[i]
						resultStack[i+1] = 1
						// fmt.Printf("merge %v for %v\n", resultStack[i], 2**16)
					}
				}
			}
		}

	}
	// brutal merge

	var result int64
	result = 1
	for _, n := range resultStack {
		result = result * n
	}
	// fmt.Printf("big merge %v from \n%v/%v\n", result, resultStack, dividStack)
	return result
}

func testDivide(ar []int) bool {
	for _, e := range ar {
		if e != 1 {
			return true
		}
	}
	return false
}
