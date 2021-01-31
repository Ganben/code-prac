package LongestConseqSeques

func longestConsecutive(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}
	longestStreak := 0
	for num := range numSet {
		if !numSet[num-1] {
			currentNum := num
			currentStrek := 1
			for numSet[currentNum+1] {
				currentNum++
				currentStrek++
			}
			if longestStreak < currentStrek {
				longestStreak = currentStrek
			}
		}
	}
	return longestStreak
}
