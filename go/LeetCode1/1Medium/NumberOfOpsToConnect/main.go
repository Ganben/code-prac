package main

import "fmt"

func makeConnected(n int, connections [][]int) int {
	ans := 0
	if len(connections) < n-1 {
		return -1
	}
	graph := make([][]int, n)
	for _, c := range connections {
		v, w := c[0], c[1]
		graph[v] = append(graph[v], w)
		graph[w] = append(graph[w], v)
	}

	vis := make([]bool, n)
	var dfs func(int)
	dfs = func(from int) {
		vis[from] = true
		for _, to := range graph[from] {
			if !vis[to] {
				dfs(to)
			}
		}
	}
	//
	for i, v := range vis {
		if !v {
			ans++
			dfs(i)
		}
	}
	return ans - 1
}

func main() {
	n := 6
	connections := [][]int{
		{0, 1},
		{0, 2},
		{0, 3},
		{1, 2},
		{1, 3},
	}
	fmt.Printf("%v", makeConnected(n, connections))
}
