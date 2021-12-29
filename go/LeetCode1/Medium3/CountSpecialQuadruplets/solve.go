package countspecialquadruplets

func countspecialquadruplets(nums []int) int {
	ans := 0
	for a, x := range nums {
		for b := a + 1; b < len(nums); b++ {
			for c := b + 1; c < len(nums); c++ {
				for _, y := range nums[c+1:] {
					if x+nums[b]+nums[c] == y {
						ans++
					}
				}

			}
		}
	}
	return ans
}
