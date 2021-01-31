package solution

func singleNumbers(nums []int) []int {
	onetime := make([]int, len(nums))

	for i, n := range nums {
		for _, k := range nums {
			if k == n {
				onetime[i]++
			}
		}
	}
	res := []int{}
	for i, n := range onetime {
		if n == 1 {
			res = append(res, nums[i])
		}
	}
	return res
}
