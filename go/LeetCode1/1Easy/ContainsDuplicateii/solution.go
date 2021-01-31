package ContainsDuplicateii

func containsNearbyDuplicate(nums []int, k int) bool {
	hmap := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		if _, ok := hmap[nums[i]]; ok {
			return true
		}
		hmap[nums[i]] = true
		if len(hmap) > k {
			delete(hmap, nums[i-k])
		}
	}
	return false
}
