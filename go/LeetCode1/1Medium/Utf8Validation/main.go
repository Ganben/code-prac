package main

import "fmt"

func validUtf8(data []int) bool {
	nums := 0
	mask1 := 1 << 7
	mask2 := 1 << 6
	for i := 0; i < len(data); i++ {
		if nums == 0 {
			mask := 1 << 7
			for mask&data[i] != 0 {
				nums += 1
				mask = mask >> 1
			}
			//
			if nums == 0 {
				continue
			}

			if nums > 4 || nums == 1 {
				return false
			}
		} else {
			if !((data[i]&mask1) != 0 && (mask2&data[i]) == 0) {
				return false
			}
		}
		nums--
	}
	return nums == 0
}

func main() {
	arr := []int{197, 130, 1}
	fmt.Printf("%v\n", validUtf8(arr))
}
