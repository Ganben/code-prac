package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type hp struct {
	sort.IntSlice
	size int
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func (h *hp) push(v int) {
	h.size++
	heap.Push(h, v)
}

func (h *hp) pop() int {
	h.size--
	return heap.Pop(h).(int)
}

func (h *hp) prune() {
	for h.Len() > 0 {
		num := h.IntSlice[0]
		if h == small {
			num = -num
		}
		if d, has := delayed[num]; has {
			if d > 1 {
				delayed[num]--
			} else {
				delete(delayed, num)
			}
			heap.Pop(h)
		} else {
			break
		}
	}
}

var delayed map[int]int
var small, large *hp

func medianSlidingWindow(nums []int, k int) []float64 {
	delayed = map[int]int{}
	small = &hp{}
	large = &hp{}
	makeBalance := func() {
		if small.size > large.size+1 {
			large.push(-small.pop())
			small.prune()
		} else if small.size < large.size {
			small.push(-large.pop())
			large.prune()
		}
	}
	insert := func(num int) {
		if small.Len() == 0 || num <= -small.IntSlice[0] {
			small.push(-num)
		} else {
			large.push(num)
		}
		makeBalance()
	}
	erase := func(num int) {
		delayed[num]++
		if num <= -small.IntSlice[0] {
			small.size--
			if num == -small.IntSlice[0] {
				small.prune()
			}
		} else {
			large.size--
			if num == large.IntSlice[0] {
				large.prune()
			}
		}
		makeBalance()
	}
	getMedian := func() float64 {
		if k&1 > 0 {
			return float64(-small.IntSlice[0])
		}
		return float64(-small.IntSlice[0]+large.IntSlice[0]) / 2
	}
	for _, num := range nums[:k] {
		insert(num)
	}
	n := len(nums)
	ans := make([]float64, 1, n-k+1)
	ans[0] = getMedian()
	for i := k; i < n; i++ {
		insert(nums[i])
		erase(nums[i-k])
		ans = append(ans, getMedian())
	}
	return ans
}

func main() {
	arr := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Printf("%v\n", medianSlidingWindow(arr, k))
}
