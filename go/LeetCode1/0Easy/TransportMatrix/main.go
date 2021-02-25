package main

import "fmt"

func transpose(matrix [][]int) [][]int {
	n, m := len(matrix), len(matrix[0])
	t := make([][]int, m)
	for i := range t {
		t[i] = make([]int, n)
		for j := range t[i] {
			t[i][j] = -1
		}
	}
	for i, row := range matrix {
		for j, v := range row {
			t[j][i] = v
		}
	}
	return t
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Printf("%v\n", transpose(matrix))
}
