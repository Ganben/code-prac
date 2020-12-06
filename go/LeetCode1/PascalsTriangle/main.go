package main

import "fmt"

func generate(numRows int) [][]int {
	var res [][]int
	for i := 0; i < numRows; i++ {
		rows := []int{}
		tmp := make([]int, i+1)
		if i == 0 {
			rows = []int{1}
		} else {
			if i%2 == 1 {
				// even row
				// length = i + 1
				for j := 0; j < (i+1)/2; j++ {
					if j == 0 {
						rows = append(rows, 1)
					}
					if j > 0 {
						rows = append(rows, res[i-1][j-1]+res[i-1][j])
					}
				}
				for j := (i + 1) / 2; j < i+1; j++ {
					rows = append(rows, rows[i-j])
				}
			} else {
				// odd row
				// length = i + 1, half length = i / 2
				for j := 0; j <= i/2; j++ {
					if j == 0 {
						rows = append(rows, 1)
					}
					if j > 0 {
						rows = append(rows, res[i-1][j-1]+res[i-1][j])
					}
				}
				//
				for j := i/2 + 1; j < i+1; j++ {
					rows = append(rows, rows[i-j])
				}
			}
			//

		}
		copy(tmp, rows)
		res = append(res, tmp)
	}
	return res
}

func main() {
	n := 5
	res := generate(n)
	fmt.Printf("%v", res)
}
