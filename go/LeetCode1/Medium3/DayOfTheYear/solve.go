package dayoftheyear

import "strconv"

func dayOfYear(date string) int {
	hmap := map[int]int{}
	hmap[1] = 31
	hmap[2] = 28
	hmap[3] = 31
	hmap[4] = 30
	hmap[5] = 31
	hmap[6] = 30
	hmap[7] = 31
	hmap[8] = 31
	hmap[9] = 30
	hmap[10] = 31
	hmap[11] = 30
	hmap[12] = 31
	//test luna
	year, _ := strconv.Atoi(date[:4])
	month, _ := strconv.Atoi(date[5:7])
	day, _ := strconv.Atoi(date[8:])
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		hmap[2] += 1
	}
	ans := day
	for i := 1; i < month; i++ {
		ans += hmap[i]
	}
	return ans

}
