package main

import "fmt"

func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) < 2 {
		return false
	}
	if len(coordinates) == 2 {
		return true
	}
	//
	k1, k2 := 0, 0
	for i, d := range coordinates {
		if i == 0 {
			continue
		}
		//
		if i == 1 {
			k1 = d[1] - coordinates[0][1]
			k2 = d[0] - coordinates[0][0]
			continue
		}
		//
		if k1*(d[0]-coordinates[0][0]) != k2*(d[1]-coordinates[0][1]) {
			return false
		}
	}
	return true

}

func main() {
	i := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 6},
		{6, 7},
	}
	fmt.Printf("%v", checkStraightLine(i))
}
