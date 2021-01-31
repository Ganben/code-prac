package ContestW

func numTriplets(nums1 []int, nums2 []int) int {
	n1 := checker(nums1, nums2)
	n2 := checker(nums2, nums1)
	return n1 + n2
}

func checker(nums1 []int, nums2 []int) int {
	sq := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		sq[i] = nums1[i] * nums1[i]
	}
	//
	ans := 0
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2)-1; j++ {
			for k := j + 1; k < len(nums2); k++ {
				if sq[i] == nums2[j]*nums2[k] {
					ans++
				}
			}
		}
	}
	return ans
}
