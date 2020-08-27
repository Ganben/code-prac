package MajorityElementii

func majorityElement(nums []int) []int {
	a, b, countA, countB := 0, 0, 0, 0
	res := []int{}
	for _, i := range nums {
		if a == i {
			countA++
			continue
		}

		if b == i {
			countB++
			continue
		}

		if countA == 0 {
			a = i
			countA = 1
			continue
		}

		if countB == 0 {
			b = i
			countB = 1
			continue
		}
		countA--
		countB--

	}

	countA, countB = 0, 0
	for _, n := range nums {
		if n == a {
			countA++
		} else if n == b {
			countB++
		}
	}

	if countA > len(nums)/3 {
		res = append(res, a)
	}
	if countB > len(nums)/3 {
		res = append(res, b)
	}
	return res

}
