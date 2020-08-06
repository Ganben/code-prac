package RotateArr

func rotate(nums []int, k int) {
	n := len(nums)
	if k > n {
		k = k % n
	}
	ans := []int{}
	ans = append(ans, nums[n-k:]...)
	ans = append(ans, nums[0:n-k]...)
	for i, v := range ans {
		nums[i] = v
	}
}
