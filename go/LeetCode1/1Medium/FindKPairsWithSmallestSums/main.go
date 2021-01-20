package main

import (
	"container/heap"
	"fmt"
)

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	h := IntHeap{}
	if len(nums2) == 0 || len(nums1) == 0 {
		return [][]int{}
	}
	//init
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			curList := []int{nums1[i], nums2[j]}
			heap.Push(&h, curList)
		}
	}
	res := [][]int{}
	for k > 0 {
		if len(h) == 0 {
			return res
		}
		u := heap.Pop(&h).([]int)
		res = append(res, u)
		k--
	}
	return res

}

type IntHeap [][]int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i][0]+h[i][1] < h[j][0]+h[j][1] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *IntHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func main() {
	u1 := []int{1, 7, 11}
	u2 := []int{2, 4, 6}
	k := 3
	fmt.Printf("%v", kSmallestPairs(u1, u2, k))
}
