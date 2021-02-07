package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

func isRectangleCover(rectangles [][]int) bool {
	hmap := map[Point]int{}
	suma := 0
	for i := 0; i < len(rectangles); i++ {
		rec := rectangles[i]
		suma += area(rec)
		p1, p2 := Point{rec[0], rec[1]}, Point{rec[2], rec[3]}
		_, has1 := hmap[p1]
		_, has2 := hmap[p2]
		if has1 {
			hmap[p1] += 1
		} else {
			hmap[p1] = 1
		}
		if has2 {
			hmap[p2] += 1
		} else {
			hmap[p2] = 1
		}
		//
		p3, p4 := Point{rec[0], rec[3]}, Point{rec[2], rec[1]}
		_, has3 := hmap[p3]
		_, has4 := hmap[p4]
		if has3 {
			hmap[p3] += 1
		} else {
			hmap[p3] = 1
		}
		if has4 {
			hmap[p4] += 1
		} else {
			hmap[p4] = 1
		}
	}
	//
	finalps := []Point{}
	for k, v := range hmap {
		if v == 1 {
			finalps = append(finalps, k)
			// fmt.Printf("%v-%v\n", k.x, k.y)
		}
	}
	//
	minX, minY := math.MaxInt32, math.MaxInt32
	maxX, maxY := 1-math.MaxInt32, 1-math.MaxInt32
	for _, p := range finalps {
		if minX > p.x {
			minX = p.x
		}
		if minY > p.y {
			minY = p.y
		}
		if maxX < p.x {
			maxX = p.x
		}
		if maxY < p.y {
			maxY = p.y
		}
		//
	}
	a1, a2, a3, a4 := 0, 0, 0, 0
	c := 0
	for _, p := range finalps {
		if p.x == minX && p.y == minY {
			a1 += 1
			c++
		}
		if p.x == minX && p.y == maxY {
			a2 += 1
			c++
		}
		if p.x == maxX && p.y == minY {
			a3 += 1
			c++
		}
		if p.x == maxX && p.y == maxY {
			a4 += 1
			c++
		}
		//

	}
	if c != 4 || a1 != 1 || a2 != 1 || a3 != 1 || a4 != 1 {
		fmt.Printf("c%v,%v,%v,%v,%v\n", c, a1, a2, a3, a4)
		return false
	}
	//area test
	if suma != area([]int{minX, minY, maxX, maxY}) {
		fmt.Printf("a%v\n", suma)
		return false
	}
	//point test
	if len(finalps) != 4 {
		fmt.Printf("%v\n", finalps)
		return false
	}
	return true

}

func area(rec []int) int {
	// area of given rec
	return (rec[2] - rec[0]) * (rec[3] - rec[1])
}

func main() {
	recs := [][]int{
		{1, 1, 3, 3},
		{3, 1, 4, 2},
		{3, 2, 4, 4},
		{1, 3, 2, 4},
		{2, 3, 3, 4},
	}
	fmt.Printf("%v\n", isRectangleCover(recs))
}
