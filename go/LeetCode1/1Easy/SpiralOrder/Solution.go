package Solution

var ans []int

func spiralOrder(matrix [][]int) []int {
	ans = []int{}
	walker(matrix, 0)
	return ans
}

func walker(matrix [][]int, stack int) {
	i, j := 0, stack
	d := 0
	n := len(matrix)
	if n == 0 {
		return
	}
	m := len(matrix[0]) - 2*stack
	if m == 0 {
		return
	}
	if n == 1 {
		p := matrix[0]
		for _, n := range p[stack : stack+m] {
			ans = append(ans, n)
		}
		return
	}
	for inc := 0; inc < (2*m + 2*n - 3); inc++ {
		switch d {
		case 0:
			// east
			if j+1 < stack+m {
				// fmt.Printf("add %v, at %v,%v\n", matrix[i][j], i, j)
				ans = append(ans, matrix[i][j])
				j++
			} else {
				// fmt.Printf("add %v, at %v,%v\n\n", matrix[i][j], i, j)
				ans = append(ans, matrix[i][j])
				i++
				d = 1
			}
		case 1:
			// south
			if i+1 < n {
				// fmt.Printf("add %v, at %v,%v\n", matrix[i][j], i, j)
				ans = append(ans, matrix[i][j])
				i++
			} else {
				// fmt.Printf("add %v, at %v,%v\n\n", matrix[i][j], i, j)
				ans = append(ans, matrix[i][j])
				if j > stack {
					j--
					d = 2
				} else {
					return
				}

			}
		case 2:
			// west
			if j > stack {
				// fmt.Printf("add %v, at %v,%v\n", matrix[i][j], i, j)
				ans = append(ans, matrix[i][j])
				j--
			} else {
				// fmt.Printf("add %v, at %v,%v\n\n", matrix[i][j], i, j)
				ans = append(ans, matrix[i][j])
				i--
				d = 3
			}
		case 3:
			if i > 0 {
				// fmt.Printf("add %v, at %v,%v\n", matrix[i][j], i, j)
				ans = append(ans, matrix[i][j])
				i--
			} else {

				// fmt.Printf("deeper %v\n\n", matrix[1:n-1])

				walker(matrix[1:n-1], stack+1)
			}
		}
	}
}
