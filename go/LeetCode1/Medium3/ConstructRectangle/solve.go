package constructrectangle

import "math"

func constructRectangle(area int) []int {
	w := 1
	// l := area / w
	res := []int{}
	w = int(math.Sqrt(float64(area)))
	for w >= 1 {
		if area%w == 0 {
			res = append(res, area/w)
			res = append(res, w)
			// return res
		}
		w -= 1
	}
	return res[:2]
}
