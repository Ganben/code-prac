package solution

func merge(nums1 []int, m int, nums2 []int, n int) {
	num := make([]int, m+n)
	p1, p2 := 0, 0
	for i := 0; i < m+n; i++ {
		if p1 < m && p2 < n {
			if nums1[p1] <= nums2[p2] {
				num[i] = nums1[p1]
				p1++
			} else {
				num[i] = nums2[p2]
				p2++
			}
		} else {
			if p1 < m {
				num = append(num, nums1[p1:]...)
				break
			} else {
				num = append(num, nums2[p2:]...)
				break
			}
		}
	}
	nums1 = num
}
