package randomflipmatrix

import "math/rand"

type Solution struct {
	m, n, total int
	mp          map[int]int
}

func Constructor(m int, n int) Solution {
	return Solution{m, n, m * n, map[int]int{}}
}

func (this *Solution) Flip() []int {
	var ans []int
	x := rand.Intn(this.total)
	this.total--
	if y, ok := this.mp[x]; ok {
		ans = []int{y / this.n, y % this.n}
	} else {
		ans = []int{x / this.n, x % this.n}
	}
	if y, ok := this.mp[this.total]; ok {
		this.mp[x] = y
	} else {
		this.mp[x] = this.total
	}
	return ans
}

func (this *Solution) Reset() {
	this.total = this.m * this.n
	this.mp = map[int]int{}
}
