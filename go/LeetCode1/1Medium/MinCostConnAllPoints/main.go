package main

import (
	"fmt"
	"math"
	"sort"
)

type unionFind struct {
	parent, rank []int
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
		rank[i] = 1
	}
	return &unionFind{parent, rank}
}

func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *unionFind) union(x, y int) bool {
	fx, fy := uf.find(x), uf.find(y)
	if fx == fy {
		return false
	}

	if uf.rank[fx] < uf.rank[fy] {
		fx, fy = fy, fx
	}
	uf.rank[fx] += uf.rank[fy]
	uf.parent[fy] = fx
	return true
}

type fenwickTree struct {
	tree, idRec []int
}

func newFenwickTree(n int) *fenwickTree {
	tree := make([]int, n)
	idRec := make([]int, n)
	for i := range tree {
		tree[i], idRec[i] = math.MaxInt64, -1
	}
	return &fenwickTree{tree, idRec}
}

func (f *fenwickTree) update(pos, val, id int) {
	for ; pos > 0; pos &= pos - 1 {
		if val < f.tree[pos] {
			f.tree[pos], f.idRec[pos] = val, id
		}
	}
}

func (f *fenwickTree) query(pos int) int {
	minVal, minID := math.MaxInt64, -1
	for ; pos < len(f.tree); pos += pos & -pos {
		if f.tree[pos] < minVal {
			minVal, minID = f.tree[pos], f.idRec[pos]
		}
	}
	return minID
}

func dist(p, q []int) int {
	return abs(p[0]-q[0]) + abs(p[1]-q[1])
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minCostConnectPoints(points [][]int) int {

	ans := 0
	n := len(points)
	for i, p := range points {
		points[i] = append(p, i)
	}
	type edge struct{ v, w, dis int }
	edges := []edge{}

	build := func() {
		sort.Slice(points, func(i, j int) bool {
			a, b := points[i], points[j]
			return a[0] < b[0] || a[0] == b[0] && a[1] < b[1]
		})
		//
		type pair struct{ v, i int }
		ps := make([]pair, n)
		for i, p := range points {
			ps[i] = pair{p[1] - p[0], i}
		}
		//
		sort.Slice(ps, func(i, j int) bool {
			return ps[i].v < ps[j].v
		})
		//
		kth := make([]int, n)
		k := 1
		kth[ps[0].i] = k
		for i := 1; i < n; i++ {
			if ps[i].v != ps[i-1].v {
				k++
			}
			kth[ps[i].i] = k
		}
		//
		t := newFenwickTree(k + 1)
		for i := n - 1; i >= 0; i-- {
			p := points[i]
			pos := kth[i]
			if j := t.query(pos); j != -1 {
				q := points[j]
				edges = append(edges, edge{p[2], q[2], dist(p, q)})
			}
			t.update(pos, p[0]+p[1], i)
		}
	}
	//
	build()
	for _, p := range points {
		p[0], p[1] = p[1], p[0]
	}
	build()
	for _, p := range points {
		p[0] = -p[0]
	}
	build()
	for _, p := range points {
		p[0], p[1] = p[1], p[0]
	}
	build()
	//
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dis < edges[j].dis
	})
	//
	uf := newUnionFind(n)
	left := n - 1
	for _, e := range edges {
		if uf.union(e.v, e.w) {
			ans += e.dis
			left--
			if left == 0 {
				break
			}
		}
	}
	return ans
}

func main() {
	p := [][]int{
		{3, 12},
		{-2, 5},
		{-4, 1},
	}
	fmt.Printf("%v", minCostConnectPoints(p))
}
