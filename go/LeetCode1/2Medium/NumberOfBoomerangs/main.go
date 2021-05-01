package main

import "fmt"

func numberOfBoomerangs(points [][]int) int {
	ans := 0
	hash := map[int]int{}
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points); j++ {
			if i != j {
				d := dist(points[i], points[j])
				if hash[d] > 0 {
					ans += hash[d] * 2
				}
				hash[d]++
			}
		}
		hash = map[int]int{}
	}
	return ans
}

func dist(i, j []int) int {
	return (i[0]-j[0])*(i[0]-j[0]) + (i[1]-j[1])*(i[1]-j[1])
}

func main() {
	i := [][]int{
		{0, 0},
		{1, 0},
		{2, 0},
	}
	//
	fmt.Printf("%v\n", numberOfBoomerangs(i))
}
