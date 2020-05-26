package Solution

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	start := 0
	end := len(nums) - 1
	mid := 0
	for start <= end {
		mid = start + int((end-start)/2)
		if nums[mid] == target {
			return true
		}
		if nums[start] == nums[mid] {
			start++
			continue
		}
		if nums[start] < nums[mid] {
			if nums[mid] > target && nums[start] <= target {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if nums[mid] < target && nums[end] >= target {
				start = mid + 1
			} else {
				end = mid - 1
			}

		}
	}
	return false
}
