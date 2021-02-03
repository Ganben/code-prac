package main

import "fmt"

var list []int

func lexicalOrder(n int) []int {
	list = []int{}
	for i := 1; i < 10; i++ {
		dfs(n, i)
	}
	return list
}

func dfs(n, i int) {
	if i > n {
		return
	}
	list = append(list, i)
	for j := 0; j <= 9; j++ {
		dfs(n, i*10+j)
	}
}

func main() {
	n := 13
	fmt.Printf("%v\n", lexicalOrder(n))
}
