package CountOfRangeSum

import "math"

func countRangeSum1(nums []int, lower int, upper int) int {
	var mergeCount func([]int) int
	mergeCount = func(arr []int) int {
		n := len(arr)
		if n <= 1 {
			return 0
		}
		n1 := append([]int(nil), arr[:n/2]...)
		n2 := append([]int(nil), arr[n/2:]...)
		cnt := mergeCount(n1) + mergeCount(n2) //
		l, r := 0, 0
		for _, v := range n1 {
			for l < len(n2) && n2[l]-v < lower {
				l++
			}
			for r < len(n2) && n2[r]-v <= upper {
				r++
			}
			cnt += r - l
		}

		p1, p2 := 0, 0
		for i := range arr {
			if p1 < len(n1) && (p2 == len(n2)) || n1[p1] <= n2[p2] {
				arr[i] = n1[p1]
				p1++
			} else {
				arr[i] = n2[p2]
				p2++
			}
		}
		return cnt
	}
	prefixSum := make([]int, len(nums)+1)
	for i, v := range nums {
		prefixSum[i+1] = prefixSum[i] + v
	}
	return mergeCount(prefixSum)
}

type node struct {
	l, r, val int
	lo, ro    *node
}

func (o *node) insert(val int) {
	o.val++
	if o.l == o.r {
		return
	}
	m := (o.l + o.r) >> 1
	if val <= m {
		if o.lo == nil {
			o.lo = &node{l: o.l, r: m}
		}
		o.lo.insert(val)
	} else {
		if o.ro == nil {
			o.ro = &node{l: m + 1, r: o.r}
		}
		o.ro.insert(val)
	}
}

func (o *node) query(l, r int) (res int) {
	if o == nil || l > o.r || r < o.l {
		return
	}
	if l <= o.l && o.r <= r {
		return o.val
	}
	return o.lo.query(l, r) + o.ro.query(l, r)
}

func countRangeSum(nums []int, lower int, upper int) (cnt int) {
	preSum := make([]int, len(nums)+1)
	for i, v := range nums {
		preSum[i+1] = preSum[i] + v
	}

	lbound, rbound := math.MaxInt64, -math.MaxInt64
	for _, sum := range preSum {
		lbound = min(lbound, sum, sum-lower, sum-upper)
		rbound = max(rbound, sum, sum-lower, sum-upper)

	}
	root := &node{l: lbound, r: rbound}
	for _, sum := range preSum {
		left, right := sum-upper, sum-lower
		cnt += root.query(left, right)
		root.insert(sum)
	}
	return
}

func min(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v < res {
			res = v
		}
	}
	return res
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}
