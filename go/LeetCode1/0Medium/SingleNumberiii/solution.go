package SingleNumberiii

func singleNumber(nums []int) []int {
	nnmap := map[int]int{}
	for _, n := range nums {
		if _, ok := nnmap[n]; ok {
			nnmap[n]++
		} else {
			nnmap[n] = 1
		}
	}
	ans := []int{}
	for _, n := range nums {
		if nnmap[n] == 1 {
			ans = append(ans, n)
		}
	}
	return ans
}

func xorSingle(nums []int) []int {
	var result []int
	if nums == nil {
		return result
	}
	xor := 0
	for _, v := range nums {
		xor ^= v
	}
	position := 0
	for i := xor; i&1 == 0; i >>= 1 {
		position++
	}
	// result is one of the two single nums
	// xor result xor again with position all 1 nums
	tempXor := xor
	for _, v := range nums {
		if (uint(v)>>uint(position))&1 == 1 {
			xor ^= v
		}
	}
	// get another xor by xor back

	return []int{xor, xor ^ tempXor}
}
