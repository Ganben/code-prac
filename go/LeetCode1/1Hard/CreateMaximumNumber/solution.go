package CreateMaximumNumber

import "fmt"

var ans [][]int

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	// cur1, cur2 := 0, 0
	ans = [][]int{}
	// recursive
	cur := []int{}
	bfs(cur, nums1, nums2, k)
	// vs := make([]int, len(ans))
	maxval := make([]int, k)
	maxid := 0
	for i, value := range ans {
		if compArr(value, maxval) {
			maxid = i
		}
	}
	return ans[maxid]
}

func compArr(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for idx, ne := range arr1 {
		if ne < arr2[idx] {
			return false
		}
		if ne > arr2[idx] {
			return true
		}
	}
	return false
}

func bindArr(arr []int) int {
	ans := 0
	multiplier := 1
	for i := len(arr) - 1; i >= 0; i-- {
		ans += multiplier * arr[i]
		multiplier *= 10
	}
	return ans
}

func bfs(cur, nums1, nums2 []int, k int) {
	if len(cur) == k {

		ans = append(ans, cur)
		return
	}

	if len(nums1)+len(nums2) >= k {
		// slice
		currentMax1, newn11, newn12 := maxArrayRestLeft(nums1, nums2, k-len(cur))

		currentMax2, newn21, newn22 := maxArrayRestRight(nums1, nums2, k-len(cur))

		if currentMax1 > currentMax2 {
			cur = append(cur, currentMax1)
			bfs(cur, newn11, newn12, k)
		} else if currentMax1 < currentMax2 {
			cur = append(cur, currentMax2)
			bfs(cur, newn21, newn22, k)
		} else {
			// equal
			// equal
			cur1 := append(cur, currentMax1)
			bfs(cur1, newn11, newn12, k)
			if len(newn11) != len(newn21) {
				// cur2 := append(cur, currentMax2)
				fmt.Printf("new route with %v - %v\n", newn11, newn21)
				bfs(cur1, newn21, newn22, k)
			}
		}

	}
}

func maxArrayRestLeft(n1, n2 []int, k int) (val int, newn1, newn2 []int) {
	p := len(n1) + len(n2) - k

	val = 0
	idx := 0
	nx := append(n1, n2...)
	for i, value := range nx {
		if i <= p {
			if val < value {
				val = value
				idx = i
			}
		} else {
			break
		}

	}

	if idx > len(n1)-1 {
		newn1 = n1
		newn2 = n2[idx-len(n1)+1:]
	} else {
		newn1 = n1[idx+1:]
		newn2 = n2
	}
	return val, newn1, newn2
}

func maxArrayRestRight(n1, n2 []int, k int) (val int, newn1, newn2 []int) {
	p := len(n1) + len(n2) - k

	val = 0
	idx := 0
	nx := append(n2, n1...)
	for i, value := range nx {
		if i <= p {
			if val < value {
				val = value
				idx = i
			}
		} else {
			break
		}
	}

	if idx >= len(n2)-1 {
		newn2 = n2
		newn1 = n1[idx-len(n2)+1:]
	} else {
		newn2 = n2[idx+1:]
		newn1 = n1
	}
	return val, newn1, newn2
}

func maxArray(n1, n2 []int) (val int, newn1 []int, newn2 []int) {
	val = 0
	idx := 0
	choice := 0
	for i, v := range n1 {
		if v > val {
			val = v
			idx = i
			choice = 1
		}
	}
	for i, value := range n2 {
		if value > val {
			val = value
			idx = i
			choice = 2
		}
	}
	if choice == 1 {
		if idx < len(n1)-1 {
			newn1 = n1[idx+1:]
		} else {
			newn1 = []int{}
		}
		newn2 = n2
	} else if choice == 2 {
		if idx < len(n2)-1 {
			newn2 = n2[idx+1:]
		} else {
			newn2 = []int{}
		}
		newn1 = n1
	}
	return val, newn1, newn2
}
