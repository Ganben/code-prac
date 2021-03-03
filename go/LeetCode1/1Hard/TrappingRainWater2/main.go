package main

import (
	"container/list"
	"fmt"
	"sort"
)

type RainPoint struct {
	x, y, v   int
	IsInQueue bool
}

type RainPointList []*RainPoint

func maxi(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (r RainPointList) Len() int {
	return len(r)
}

func (r RainPointList) Less(i, j int) bool {
	return r[i].v < r[j].v
}

func (r RainPointList) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func trapRainWater(height [][]int) int {
	if len(height) < 3 || len(height[0]) < 3 {
		return 0
	}
	//
	rtn := 0
	roundList := make(RainPointList, 0)
	m := map[int]map[int]*RainPoint{}
	xl := len(height[0])
	yl := len(height)
	for y, lines := range height {
		for x, v := range lines {
			rp := &RainPoint{
				x:         x,
				y:         y,
				v:         v,
				IsInQueue: false,
			}
			if x == 0 || y == 0 || y == len(height)-1 || x == len(lines)-1 {
				rp.IsInQueue = true
				roundList = append(roundList, rp)
			}
			if m[y] == nil {
				m[y] = make(map[int]*RainPoint)
			}
			m[y][x] = rp
		}
	}
	//
	sort.Sort(roundList)
	l := list.New()
	for _, v := range roundList {
		l.PushBack(v)
	}
	rm := [][]int{
		{0, -1},
		{0, 1},
		{-1, 0},
		{1, 0},
	}
	maxp := 0
	for element := l.Front(); element != nil; element = element.Next() {
		//
		rp := element.Value.(*RainPoint)
		maxp = maxi(rp.v, maxp)
		for _, v := range rm {
			x := rp.x + v[0]
			y := rp.y + v[1]
			if x < 0 || x >= xl || y < 0 || y >= yl {
				continue
			}
			p := m[y][x]
			if p.IsInQueue {
				continue
			}
			if p.v < maxp {
				rtn += maxp - p.v

			}
			//
			el := element.Next()
			for {
				if el == nil {
					l.PushBack(p)
					break
				}
				if p.v <= el.Value.(*RainPoint).v {
					l.InsertBefore(p, el)
					break
				}
				el = el.Next()
			}
			p.IsInQueue = true
		}
	}
	return rtn
}

func main() {
	harrs := [][]int{
		{1, 4, 3, 1, 3, 2},
		{3, 2, 1, 3, 2, 4},
		{2, 3, 3, 2, 3, 1},
	}
	//
	fmt.Printf("%v\n", trapRainWater(harrs))
}
