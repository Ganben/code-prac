package main

import "fmt"

func wiggleSort(nums []int) {
	wiggleSortRe(nums, 0, len(nums)-1, 0, len(nums)-1, nums[0])
}

func wiggleSortRe(nums []int, s, e, l, r, q int) {
	// s := 0
	// e := len(nums) - 1
	// l, r := s, e
	// q := nums[s]
	if s > e {
		return
	}
	for l < r {
		for l < r && nums[r] >= q {
			r--
		}
		for l < r && nums[l] <= q {
			l++
		}
		//swap
		nums[l], nums[r] = nums[r], nums[l]

	}
	nums[l], nums[r] = nums[r], nums[l]
	if l > len(nums)>>1 {
		wiggleSortRe(nums, s, l-1, s, l-1, nums[s])
	} else if l < len(nums)>>1 {
		wiggleSortRe(nums, l+1, e, l+1, e, nums[l+1])
	}
	//
	q = nums[l]
	l = -1
	r = len(nums)
	for i := 0; i < r; i++ {

		for nums[i] > q && i < r {
			r--
			nums[i], nums[r] = nums[r], nums[i]

			if r < 0 {
				break
			}
		}
		for nums[i] < q && i > l {
			l++
			nums[i], nums[l] = nums[l], nums[i]

			if l > len(nums)-1 {
				break
			}
		}
	}
	l = (len(nums) + 1) >> 1
	r = len(nums)
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	for i, _ := range nums {
		if i&1 != 0 {
			r--
			nums[i] = tmp[r]
		} else {
			l--
			nums[i] = tmp[l]
		}

	}
}

func main() {
	nums := []int{1, 5, 1, 1, 6, 4}
	wiggleSort(nums)
	fmt.Printf("%v", nums)
}
