func wiggleSort(nums []int) {

}

func threeWayPartition(nums []int, left, right int) {
	// use golang's slice ref, like &a[]int{}
	lt := left
	i := left
	gt := right
	partitionValue := nums[left]
	for i <= gt {
		if nums[i] < partitionValue {
			nums[i], nums[lt] = nums[lt], nums[i]
			i++
			lt++
		} else {
			if nums[i] > partitionValue {
				nums[i], nums[gt] = nums[gt], nums[i]
				gt--
			} else {
				i++
			}
		}
	}
}
