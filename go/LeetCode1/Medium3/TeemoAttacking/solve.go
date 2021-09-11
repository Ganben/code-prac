package teemoattacking

func findPoisonedDuration(timeSeries []int, duration int) int {
	ans := 0
	for i, t := range timeSeries {
		if i < len(timeSeries)-1 {
			if t+duration > timeSeries[i+1] {
				ans += timeSeries[i+1] - t
			} else {
				ans += duration
			}
		} else {
			ans += duration
		}

	}
	return ans
}
