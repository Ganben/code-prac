package main

import "fmt"

func matrixScore(A [][]int) int {
	m, n := len(A), len(A[0])
	ans := 1 << (n - 1) * m
	for j := 1; j < n; j++ {
		ones := 0
		for _, row := range A {
			if row[j] == row[0] {
				ones++
			}
		}
		if ones < m-ones {
			ones = m - ones
		}
		ans += 1 << (n - 1 - j) * ones
	}
	return ans
}

func main() {
	inMatrix := [][]int{
		{0, 0, 1, 1},
		{1, 0, 1, 0},
		{1, 1, 0, 0},
	}
	s := matrixScore(inMatrix)
	fmt.Printf("%v", s)

}
