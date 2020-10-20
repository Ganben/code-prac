package PerfectSquares

func numSquaresBFS(n int) int {
	squareNums := []int{}
	for i := 1; i*i <= n; i++ {
		squareNums = append(squareNums, i*i)
	}
	queue := map[int]bool{}
	queue[n] = true
	level := 0
	for len(queue) > 0 {
		level++
		nextQueue := map[int]bool{}
		for remainder, _ := range queue {
			for square := range squareNums {
				if remainder == square {
					return level
				} else if remainder < square {
					break
				} else {
					nextQueue[remainder-square] = true
				}

			}
		}
		//
		queue = nextQueue
	}
	return level
}
