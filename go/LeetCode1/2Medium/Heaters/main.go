package main

import "sort"

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	res := 0
	for _, house := range houses {
		left := 0
		right := len(heaters) - 1
		houseRes := 0
		for left < right {
			middle := left + (right-left)/2
			if heaters[middle] < house {
				left = middle + 1
			} else {
				right = middle
			}
		}
		if heaters[left] == house {
			houseRes = 0
		} else if heaters[left] < house {
			houseRes = house - heaters[left]

		} else {
			if left == 0 {
				houseRes = heaters[left] - house
			} else {
				houseRes = getMin(heaters[left]-house, house-heaters[left-1])
			}
		}
		res = getMax(res, houseRes)
	}
	return res
}

func getMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func getMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}
