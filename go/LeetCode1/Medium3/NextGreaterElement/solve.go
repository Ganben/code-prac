package nextgreaterelement

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// sort.Ints(nums2)
	ans := make([]int, len(nums1))
	for i, v := range nums1 {
		ans[i] = -1
		for ij, vs := range nums2 {
			if vs == v {
				d := ij
				for j := d; j < len(nums2); j++ {
					if nums2[j] > v {
						ans[i] = nums2[j]
						break
					}
				}
				break
			}

		}

	}
	return ans
}
