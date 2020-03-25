package Solution

import "math/rand"

func mergeKLists(lists []*ListNode) *ListNode {
	nums := make([]int, 0, 10)
	for _, v := range lists {
		nums = append(nums, MarshalListNodeToSlice(v)...)
	}
	nums = QuickSort(nums)

	return MarshalSliceToListNode(nums)
}

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	median := arr[rand.Intn(len(arr))]

	low_part := make([]int, 0, len(arr))
	high_part := make([]int, 0, len(arr))
	middle_part := make([]int, 0, len(arr))

	for _, item := range arr {
		switch {
		case item < median:
			low_part = append(low_part, item)
		case item == median:
			middle_part = append(middle_part, item)
		case item > median:
			high_part = append(high_part, item)
		}
	}
	low_part = QuickSort(low_part)
	high_part = QuickSort(high_part)

	low_part = append(low_part, middle_part...)
	low_part = append(low_part, high_part...)
	return low_part
}

func MarshalListNodeToSlice(node *ListNode) []int {
	ans := make([]int, 0, 10)
	for node != nil {
		ans = append(ans, node.Val)
		node = node.Next
	}
	return ans
}

func MarshalSliceToListNode(nums []int) *ListNode {
	head := &ListNode{}
	tmp := head
	for _, v := range nums {
		tmp.Next = &ListNode{Val: v, Next: nil}
		tmp = tmp.Next
	}
	return head.Next
}
